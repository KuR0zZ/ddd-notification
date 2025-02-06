package repository

import (
	"fmt"
	"notification-service/domain/notification/entity"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Create(notification *entity.Notification) error
	GetNotSent() ([]entity.Notification, error)
	UpdateStatus(id string) error
}

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) Repository {
	return &PostgresRepository{db}
}

func (pr *PostgresRepository) Create(notification *entity.Notification) error {
	query := "INSERT INTO Notifications (email, message, type, is_sent) VALUES ($1, $2, $3, $4) RETURNING id"

	err := pr.db.QueryRow(query, notification.Email, notification.Message, notification.Type, notification.IsSent).Scan(&notification.ID)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostgresRepository) GetNotSent() ([]entity.Notification, error) {
	query := "SELECT * FROM Notifications WHERE status = $1"

	rows, err := pr.db.Query(query, "Not Sent")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []entity.Notification
	for rows.Next() {
		var notification entity.Notification

		err := rows.Scan(&notification.ID, &notification.Email, &notification.Message, &notification.Type, &notification.IsSent)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (pr *PostgresRepository) UpdateStatus(id string) error {
	query := "UPDATE Notifications SET status = $1 WHERE id = $2 AND status != $3"

	result, err := pr.db.Exec(query, "Sent", id, "Sent")
	if err != nil {
		return err
	}

	if row, _ := result.RowsAffected(); row == 0 {
		return fmt.Errorf("notification not found or already sent")
	}

	return nil
}
