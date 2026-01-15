package service

import (
	"context"

	"github.com/tmozzze/ChatUs/internal/models"
)

type ChatRepository interface {
	CreateChat(ctx context.Context, chat *models.Chat) error
	CreateMessage(ctx context.Context, msg *models.Message) error
	GetChatWithMessages(ctx context.Context, id uint, limit int) (*models.Chat, error)
}
