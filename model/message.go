package model

import "time"

type Message struct {
	ID        string    `db:"id"`
	RoomID    string    `db:"room_id"`
	Message   string    `db:"message"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
