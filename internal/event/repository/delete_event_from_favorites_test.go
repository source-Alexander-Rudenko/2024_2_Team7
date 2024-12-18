package eventRepository

import (
	"context"
	"fmt"
	"testing"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"kudago/internal/models"
)

func TestEventRepository_DeleteEventFromFavorites(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name      string
		favorite  models.FavoriteEvent
		mockSetup func(m pgxmock.PgxConnIface)
		expectErr error
	}{
		{
			name: "Успешное удаление из избранного",
			favorite: models.FavoriteEvent{
				UserID:  1,
				EventID: 2,
			},
			mockSetup: func(m pgxmock.PgxConnIface) {
				m.ExpectExec(`DELETE FROM FAVORITE_EVENT WHERE user_id=\$1 AND event_id=\$2`).
					WithArgs(1, 2).
					WillReturnResult(pgxmock.NewResult("DELETE", 1))
			},
			expectErr: nil,
		},
		{
			name: "Событие не найдено в избранном",
			favorite: models.FavoriteEvent{
				UserID:  1,
				EventID: 3,
			},
			mockSetup: func(m pgxmock.PgxConnIface) {
				m.ExpectExec(`DELETE FROM FAVORITE_EVENT WHERE user_id=\$1 AND event_id=\$2`).
					WithArgs(1, 3).
					WillReturnResult(pgxmock.NewResult("DELETE", 0))
			},
			expectErr: models.ErrNotFound,
		},
		{
			name: "Ошибка базы данных",
			favorite: models.FavoriteEvent{
				UserID:  1,
				EventID: 4,
			},
			mockSetup: func(m pgxmock.PgxConnIface) {
				m.ExpectExec(`DELETE FROM FAVORITE_EVENT WHERE user_id=\$1 AND event_id=\$2`).
					WithArgs(1, 4).
					WillReturnError(fmt.Errorf("database error"))
			},
			expectErr: fmt.Errorf("%s: database error", models.LevelDB),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockConn, err := pgxmock.NewConn()
			require.NoError(t, err)
			defer mockConn.Close(ctx)

			tt.mockSetup(mockConn)

			db := NewDB(mockConn)

			err = db.DeleteEventFromFavorites(ctx, tt.favorite)

			if tt.expectErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
