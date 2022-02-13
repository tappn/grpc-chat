package handler

import (
	"context"
	"time"

	"github.com/tappn/grpc-chat/gen/go/chat/v1"
	"github.com/tappn/grpc-chat/model"
	"github.com/tappn/grpc-chat/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Chat struct {
	mr *repository.MessageRepository
	rr *repository.RoomRepository
	chat.UnimplementedChatServiceServer
}

func NewChatHandler(msg *repository.MessageRepository, room *repository.RoomRepository) *Chat {
	return &Chat{
		mr: msg,
		rr: room,
	}
}

func (h *Chat) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*emptypb.Empty, error) {
	room, err := h.rr.FindByID(req.GetRoomId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	now := time.Now()
	c := &model.Message{
		ID:        model.CreateID(),
		RoomID:    room.ID,
		Message:   req.GetMessage(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := h.mr.Store(c); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (h *Chat) GetMessages(req *chat.GetMessagesRequest, stream chat.ChatService_GetMessagesServer) error {
	var latestID string
	for {
		mm, err := h.mr.FindAll(req.RoomId, latestID)
		if err != nil {
			return err
		}
		for _, v := range mm {
			res := &chat.GetMessagesResponse{
				Message: v.Message,
			}
			if err := stream.Send(res); err != nil {
				return err
			}
		}
		if len(mm) != 0 {
			latestID = mm[len(mm)-1].ID
		}
	}
}
