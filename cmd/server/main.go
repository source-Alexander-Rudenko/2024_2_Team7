package main

import (
	"context"
	"kudago/internal/db"
	"log"
	"net/http"

	"kudago/config"
	_ "kudago/docs"
	"kudago/internal/http/auth"
	"kudago/internal/http/events"
	"kudago/internal/middleware"
	csrfRepository "kudago/internal/repository/csrf"
	eventRepository "kudago/internal/repository/events"
	sessionRepository "kudago/internal/repository/session"
	userRepository "kudago/internal/repository/users"
	authService "kudago/internal/service/auth"
	eventService "kudago/internal/service/events"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// swag init

// @title           Swagger Vihodnoy API
// @version         1.0
// @description     This is a description of the Vihodnoy server.
// @termsOfService  http://swagger.io/terms/

func main() {
	port := config.LoadConfig()
	encryptionKey := config.LoadEncriptionKey()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Server failed to start logger: %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	pool, err := db.InitDB(sugar)
	postgresPing := pool.Ping(context.Background())
	if err != nil || postgresPing != nil {
		log.Fatalf("Failed to connect to the database")
	}
	defer pool.Close()

	redisClient, err := db.InitRedis()
	if err != nil {
		log.Fatalf("Failed to connect to the redis database")
	}

	userDB := userRepository.NewDB(pool)
	sessionDB := sessionRepository.NewDB(redisClient)
	csrfDB := csrfRepository.NewDB(redisClient)
	eventDB := eventRepository.NewDB(pool)

	authService := authService.NewService(userDB, sessionDB, csrfDB)
	eventService := eventService.NewService(eventDB)

	authHandler := auth.NewAuthHandler(&authService)
	eventHandler := events.NewEventHandler(&eventService)

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	r.HandleFunc("/session", authHandler.CheckSession).Methods("GET")
	r.HandleFunc("/profile", authHandler.Profile).Methods("GET")

	r.HandleFunc("/events/{id:[0-9]+}", eventHandler.GetEventByID).Methods("GET")
	r.HandleFunc("/events/tags/{tag}", eventHandler.GetEventsByTag).Methods("GET")
	r.HandleFunc("/events/categories/{category}", eventHandler.GetEventsByCategory).Methods("GET")
	r.HandleFunc("/events", eventHandler.GetAllEvents).Methods("GET")
	r.HandleFunc("/categories", eventHandler.GetCategories).Methods("GET")
	r.HandleFunc("/events/{id:[0-9]+}", eventHandler.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id:[0-9]+}", eventHandler.DeleteEvent).Methods("DELETE")
	r.HandleFunc("/events", eventHandler.AddEvent).Methods("POST")

	whitelist := []string{
		"/login",
		"/register",
		"/events",
		"/static",
		"/session",
		"/logout",
		"/docs",
		"/categories",
	}

	handlerWithAuth := middleware.AuthMiddleware(whitelist, authHandler, r)
	handlerWithCORS := middleware.CORSMiddleware(handlerWithAuth)
	handlerWithCSRF := middleware.CSRFMiddleware(handlerWithCORS, authHandler, encryptionKey)
	handlerWithLogging := middleware.LoggingMiddleware(handlerWithCSRF, sugar)
	handler := middleware.PanicMiddleware(handlerWithLogging)

	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
