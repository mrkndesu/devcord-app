package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/mrkndesu/devcord-app/backend/model"
)

// userRepositoryImpl は UserRepository の実装
// Firestore クライアントを保持して Firestore を操作する
type userRepositoryImpl struct {
	client *firestore.Client
}

// NewUserRepository は Firestore クライアントから UserRepository を作る
func NewUserRepository(client *firestore.Client) UserRepository {
	return &userRepositoryImpl{client: client}
}

// Create は新しいユーザーを作成し、IDや作成年月をセットする
func (r *userRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	docRef := r.client.Collection("users").NewDoc()
	user.ID = docRef.ID

	now := time.Now()
	user.CreatedYear = now.Year()
	user.CreatedMonth = int(now.Month())

	_, err := docRef.Set(ctx, map[string]interface{}{
		"handle":        user.Handle,
		"name":          user.Name,
		"email":         user.Email,
		"password":      user.Password,
		"avatar_url":    user.AvatarURL,
		"description":   user.Description,
		"birth_date":    user.BirthDate,
		"created_year":  user.CreatedYear,
		"created_month": user.CreatedMonth,
	})
	return err
}

// GetAll は全ユーザーを取得し、IDをセットして返す
func (r *userRepositoryImpl) GetAll(ctx context.Context) ([]model.User, error) {
	docs, err := r.client.Collection("users").Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var users []model.User
	for _, doc := range docs {
		var user model.User
		if err := doc.DataTo(&user); err != nil {
			return nil, err
		}
		user.ID = doc.Ref.ID
		users = append(users, user)
	}
	return users, nil
}

// GetByID は指定IDのユーザーを取得して返す
func (r *userRepositoryImpl) GetByID(ctx context.Context, id string) (*model.User, error) {
	doc, err := r.client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var user model.User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}
	user.ID = doc.Ref.ID
	return &user, nil
}

// Update は指定IDのユーザー情報を上書き保存する
func (r *userRepositoryImpl) Update(ctx context.Context, id string, updatedUser model.User) error {
	_, err := r.client.Collection("users").Doc(id).Set(ctx, updatedUser)
	return err
}

// Delete は指定IDのユーザーを削除する
func (r *userRepositoryImpl) Delete(ctx context.Context, id string) error {
	_, err := r.client.Collection("users").Doc(id).Delete(ctx)
	return err
}
