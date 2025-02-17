package userRepository

import (
	"context"
	"fmt"

	"kudago/internal/models"
)

const getSubscriptionsQuery = `
	SELECT u.id, u.username, u.email, u.url_to_avatar
	FROM "USER" u
	JOIN SUBSCRIPTION s ON s.follows_id = u.id
	WHERE s.subscriber_id = $1;
`

func (d UserDB) GetSubscriptions(ctx context.Context, ID int) ([]models.User, error) {
	rows, err := d.Pool.Query(ctx, getSubscriptionsQuery, ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", models.LevelDB, err)
	}
	defer rows.Close()

	var subscriptions []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.ImageURL)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", models.LevelDB, err)
		}
		subscriptions = append(subscriptions, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", models.LevelDB, err)
	}

	return subscriptions, nil
}
