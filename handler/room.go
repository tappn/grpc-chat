package handler

import (
	"context"
	"time"

	"github.com/tappn/grpc-chat/gen/go/room/v1"
	"github.com/tappn/grpc-chat/model"
	"github.com/tappn/grpc-chat/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Room struct {
	repo *repository.RoomRepository
	room.UnimplementedRoomServiceServer
}

func NewRoomHandler(r *repository.RoomRepository) *Room {
	return &Room{
		repo: r,
	}
}

func (r *Room) CreateRoom(ctx context.Context, req *room.CreateRoomRequest) (*room.CreateRoomResponse, error) {
	now := time.Now()
	m := &model.Room{
		ID:        model.CreateID(),
		Name:      req.GetName(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := r.repo.Store(m); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	resp := &room.CreateRoomResponse{
		Room: &room.Room{
			Id:   m.ID,
			Name: m.Name,
		},
	}
	return resp, status.Errorf(codes.OK, "ok")
}

func (r *Room) GetAllRoom(context.Context, *emptypb.Empty) (*room.GetALLRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRoom not implemented")
}

func (r *Room) GetRoom(context.Context, *room.GetRoomRequest) (*room.GetRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
