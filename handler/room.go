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
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	resp := &room.CreateRoomResponse{
		Room: &room.Room{
			Id:   m.ID,
			Name: m.Name,
		},
	}
	return resp, nil
}

func (r *Room) GetAllRoom(context.Context, *emptypb.Empty) (*room.GetALLRoomResponse, error) {
	rooms, err := r.repo.FindAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var resp room.GetALLRoomResponse
	for _, v := range rooms {
		r := &room.Room{
			Id:   v.ID,
			Name: v.Name,
		}
		resp.Rooms = append(resp.Rooms, r)
	}

	return &resp, nil
}

func (r *Room) GetRoom(ctx context.Context, req *room.GetRoomRequest) (*room.GetRoomResponse, error) {
	rooms, err := r.repo.FindByID(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := &room.GetRoomResponse{
		Room: &room.Room{
			Id:   rooms.ID,
			Name: rooms.Name,
		},
	}

	return resp, nil
}
