package session

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
)

type Session struct {
	Username string
	Token    string
	Expires  time.Time
}

type SessionDB struct {
	sessions map[string]Session
}

const ExpirationTime = 30 * time.Minute

func NewSessionDB() *SessionDB {
	return &SessionDB{
		sessions: make(map[string]Session),
	}
}

func (db *SessionDB) CreateSession(username string) Session {
	sessionToken := generateSessionToken()
	expiration := time.Now().Add(ExpirationTime)

	session := Session{
		Username: username,
		Token:    sessionToken,
		Expires:  expiration,
	}

	db.sessions[sessionToken] = session
	return session
}

func (db *SessionDB) CheckSession(r *http.Request) (*Session, bool) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return nil, false
	}

	session, exists := db.sessions[cookie.Value]
	if !exists || session.Expires.Before(time.Now()) {
		return nil, false
	}
	return &session, true
}

func generateSessionToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (db *SessionDB) DeleteSession(token string) {
	delete(db.sessions, token)
}
