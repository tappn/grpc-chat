package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tappn/grpc-chat/model"
)

type RoomRepository struct {
	db *sqlx.DB
}

func NewRoomRepository(db *sqlx.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (r *RoomRepository) Store(m *model.Room) error {
	query := `
	insert into rooms(
		id,
		name,
		created_at,
		updated_at
	) values (
		:id,
		:name,
		:created_at,
		:updated_at
	)
	`

	_, err := r.db.NamedExec(query, m)
	if err != nil {
		return err
	}

	return nil
}
