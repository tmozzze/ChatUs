package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/tmozzze/ChatUs/internal/lib/er"
	"github.com/tmozzze/ChatUs/internal/models"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) CreateChat(ctx context.Context, chat *models.Chat) error {
	const op = "repository.CreateChat"

	err := r.db.WithContext(ctx).Create(chat).Error
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *ChatRepository) CreateMessage(ctx context.Context, msg *models.Message) error {
	const op = "repository.CreateMessage"

	var chatExists int64
	err := r.db.WithContext(ctx).Model(&models.Chat{}).Where("id = ?",
		msg.ChatID).Count(&chatExists).Error

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if chatExists == 0 {
		return fmt.Errorf("%s: %w", op, er.ErrNotFound)
	}

	err = r.db.WithContext(ctx).Create(msg).Error
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *ChatRepository) GetChatWithMessages(ctx context.Context, id uint, limit int) (*models.Chat, error) {
	const op = "repository.GetChatWithMessages"

	var chat models.Chat

	err := r.db.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC").Limit(limit)
	}).First(&chat, id).Error

	if err != nil {
		// gorm error mapping
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s: %w", op, er.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, er.ErrNotFound)
	}

	return &chat, nil
}

func (r *ChatRepository) DeleteChat(ctx context.Context, id uint) error {
	const op = "repository.DeleteChat"

	result := r.db.WithContext(ctx).Delete(&models.Chat{}, id)

	if result.Error != nil {
		return fmt.Errorf("%s: %w", op, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, er.ErrNotFound)
	}

	return nil
}
