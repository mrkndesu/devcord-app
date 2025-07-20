package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/mrkndesu/devcord-app/backend/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserRepositoryFirestore は Firestore を使ったユーザーデータのリポジトリ実装
type UserRepositoryFirestore struct {
	Client *firestore.Client // Firestore クライアント
}

// Create は新しいユーザーを Firestore に作成し、IDをセットする
func (r *UserRepositoryFirestore) Create(ctx context.Context, user *model.User) error {
	docRef, _, err := r.Client.Collection("users").Add(ctx, user)
	if err != nil {
		return err
	}
	user.ID = docRef.ID
	return nil
}

// GetAll は全ユーザーを Firestore から取得する
func (r *UserRepositoryFirestore) GetAll(ctx context.Context) ([]model.User, error) {
	iter := r.Client.Collection("users").Documents(ctx)
	defer iter.Stop()

	var users []model.User
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var user model.User
		if err := doc.DataTo(&user); err != nil {
			return nil, err
		}
		user.ID = doc.Ref.ID
		users = append(users, user)
	}

	return users, nil
}

// GetByID は指定IDのユーザーを取得する
func (r *UserRepositoryFirestore) GetByID(ctx context.Context, id string) (*model.User, error) {
	doc, err := r.Client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, fmt.Errorf("not found")
		}
		return nil, err
	}

	var user model.User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}
	user.ID = doc.Ref.ID
	return &user, nil
}

// Update は指定IDのユーザー情報を Firestore で更新する
func (r *UserRepositoryFirestore) Update(ctx context.Context, id string, updatedUser model.User) error {


	_, err := r.Client.Collection("users").Doc(id).Set(ctx, map[string]interface{}{
		"handle":        updatedUser.Handle,
		"name":          updatedUser.Name,
		"email":         updatedUser.Email,    
		"password":      updatedUser.Password, 
		"avatar_url":    updatedUser.AvatarURL,
		"description":   updatedUser.Description,
		"birth_date":    updatedUser.BirthDate,
		"created_year":  updatedUser.CreatedYear,
		"created_month": updatedUser.CreatedMonth,
		"updated_at":    time.Now(), // 更新日時を記録
	}, firestore.MergeAll)
	return err
}

// Delete は指定IDのユーザーを Firestore から削除する
func (r *UserRepositoryFirestore) Delete(ctx context.Context, id string) error {
	_, err := r.Client.Collection("users").Doc(id).Delete(ctx)
	return err
}

