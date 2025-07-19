package repository

import (
	"context"

	"github.com/mrkndesu/devcord-app/backend/model"
)

// UserRepository はユーザー情報の操作を定義するインターフェース
type UserRepository interface {
	// 新規ユーザーを作成
	Create(ctx context.Context, user *model.User) error

	// 指定IDのユーザーを取得
	GetByID(ctx context.Context, userID string) (*model.User, error)

	// 全ユーザーを取得
	GetAll(ctx context.Context) ([]model.User, error)

	// 指定IDのユーザー情報を更新
	Update(ctx context.Context, userID string, updatedUser model.User) error

	// 指定IDのユーザーを削除、投稿も削除される
	Delete(ctx context.Context, userID string) error
}
