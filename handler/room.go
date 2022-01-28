package handler

import (
	"context"

	"github.com/tappn/grpc-chat/gen/go/room/v1"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Room struct {
	room.UnimplementedRoomServiceServer
}

func NewRoomHandler() *Room {
	return &Room{}
}

func (r *Room) CreateRoom(context.Context, *room.CreateRoomRequest) (*room.CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (r *Room) GetAllRoom(context.Context, *emptypb.Empty) (*room.GetALLRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRoom not implemented")
}
func (r *Room) GetRoom(context.Context, *room.GetRoomRequest) (*room.GetRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
