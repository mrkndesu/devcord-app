package repository

import (
	"context"

	"github.com/mrkndesu/devcord-app/backend/model"
)

// PostRepository は投稿データ操作のインターフェース
type PostRepository interface {
	// GetAll すべての投稿を取得
	GetAll(ctx context.Context) ([]model.Post, error)

	// Create 新しい投稿を作成
	Create(ctx context.Context, post *model.Post) error

	// GetByID 指定IDの投稿を取得
	GetByID(ctx context.Context, id string) (*model.Post, error)

	// Update 指定IDの投稿を更新
	Update(ctx context.Context, id string, updatedPost model.Post) error

	// Delete 指定IDの投稿を削除
	Delete(ctx context.Context, id string) error

	// DeleteAll すべての投稿を一括削除
	DeleteAll(ctx context.Context) error
}