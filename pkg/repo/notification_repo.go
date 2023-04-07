package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"shopper"
)

type NotificationRepo struct {
	db *sqlx.DB
}

func (r *NotificationRepo) CreateNotification(notification shopper.Notification) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, createdat, text, user_id) values ($1, $2, $3, $4) RETURNING id", notifications)
	row := r.db.QueryRow(query, notification.Title, notification.CreatedAt, notification.Text, notification.UserId)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func NewNotificationRepo(db *sqlx.DB) *NotificationRepo {
	return &NotificationRepo{
		db: db,
	}
}
