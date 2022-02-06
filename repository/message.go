package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tappn/grpc-chat/model"
)

type MessageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (r *MessageRepository) Store(c *model.Message) error {
	query := `
	insert into messages(
		id,
		room_id,
		message,
		created_at,
		updated_at
	) values (
		:id,
		:room_id,
		:message,
		:created_at,
		:updated_at
	)
	`

	_, err := r.db.NamedExec(query, c)
	if err != nil {
		return err
	}

	return nil
}
