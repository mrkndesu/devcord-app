package repository

import (
	"context"

	"github.com/mrkndesu/devcord-app/backend/model"
)

type PostRepository interface {
	// 指定ユーザーの全投稿を取得
	GetAll(ctx context.Context, userID string) ([]model.Post, error)

	// 投稿を作成
	Create(ctx context.Context, userID string, post *model.Post) error

	// 指定IDの投稿を取得
	GetByID(ctx context.Context, userID, postID string) (*model.Post, error)

	// 指定IDの投稿を更新
	Update(ctx context.Context, userID, postID string, updatedPost model.Post) error

	// 指定IDの投稿を削除
	Delete(ctx context.Context, userID, postID string) error

	// 指定ユーザーの全投稿を削除
	DeleteAll(ctx context.Context, userID string) error
}