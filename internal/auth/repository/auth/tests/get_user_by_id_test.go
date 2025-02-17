package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"
	"kudago/internal/auth/repository/auth"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"kudago/internal/models"
)

func TestUserDB_GetUserByID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name      string
		ID        int
		mockSetup func(m pgxmock.PgxConnIface)
		expected  models.User
		expectErr bool
	}{
		{
			name: "Пользователь не найден",
			ID:   2,
			mockSetup: func(m pgxmock.PgxConnIface) {
				m.ExpectQuery(`SELECT id, username, email, url_to_avatar FROM "USER" WHERE id=\$1`).
					WithArgs(2).
					WillReturnError(pgx.ErrNoRows)
			},
			expected:  models.User{},
			expectErr: true,
		},
		{
			name: "Ошибка при запросе",
			ID:   3,
			mockSetup: func(m pgxmock.PgxConnIface) {
				m.ExpectQuery(`SELECT id, username, email, url_to_avatar FROM "USER" WHERE id=\$1`).
					WithArgs(3).
					WillReturnError(fmt.Errorf("database error"))
			},
			expected:  models.User{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockConn, err := pgxmock.NewConn()
			require.NoError(t, err)
			defer mockConn.Close(ctx)

			tt.mockSetup(mockConn)

			db := userRepository.UserDB{Pool: mockConn}

			user, err := db.GetUserByID(ctx, tt.ID)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, user)
			}
		})
	}
}
