package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/mrkndesu/devcord-app/backend/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/api/iterator"
)

// PostRepositoryFirestore は Firestore を使った投稿リポジトリの実装
// users コレクションの posts サブコレクションを操作する
type PostRepositoryFirestore struct {
	Client *firestore.Client // Firestore クライアント
}

// GetAll は指定ユーザーの全投稿を取得する
func (r *PostRepositoryFirestore) GetAll(ctx context.Context, userID string) ([]model.Post, error) {
	iter := r.Client.Collection("users").Doc(userID).Collection("posts").Documents(ctx) 
	defer iter.Stop() 

	var posts []model.Post
	for {
		doc, err := iter.Next()
		if err != nil {
			break 
		}

		var post model.Post
		if err := doc.DataTo(&post); err != nil {
			continue 
		}

		post.ID = doc.Ref.ID 
		posts = append(posts, post)
	}
	return posts, nil
}

// Create は指定ユーザーの posts に新しい投稿を追加する
func (r *PostRepositoryFirestore) Create(ctx context.Context, userID string, post *model.Post) error {
	docRef, _, err := r.Client.Collection("users").Doc(userID).Collection("posts").Add(ctx, post) 
	if err != nil {
		return err
	}
	post.ID = docRef.ID 
	return nil
}

// GetByID は指定ユーザーの特定投稿を取得する
func (r *PostRepositoryFirestore) GetByID(ctx context.Context, userID, postID string) (*model.Post, error) {
	doc, err := r.Client.Collection("users").Doc(userID).Collection("posts").Doc(postID).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, fmt.Errorf("not found")
		}
		return nil, err
	}

	var post model.Post
	if err := doc.DataTo(&post); err != nil {
		return nil, err
	}
	post.ID = doc.Ref.ID
	return &post, nil
}

// Update は指定ユーザーの特定投稿を更新する
func (r *PostRepositoryFirestore) Update(ctx context.Context, userID, postID string, post model.Post) error {
	_, err := r.Client.Collection("users").Doc(userID).Collection("posts").Doc(postID).Set(ctx, map[string]interface{}{
		"user_id":    post.UserID,
		"title":      post.Title,
		"content":    post.Content,
		"updated_at": time.Now(), 
	}, firestore.MergeAll) 
	return err
}

// Delete は指定ユーザーの特定投稿を削除する
func (r *PostRepositoryFirestore) Delete(ctx context.Context, userID, postID string) error {
	_, err := r.Client.Collection("users").Doc(userID).Collection("posts").Doc(postID).Delete(ctx)
	return err
}

// DeleteAll は指定ユーザーの全投稿を削除する
func (r *PostRepositoryFirestore) DeleteAll(ctx context.Context, userID string) error {
	iter := r.Client.Collection("users").Doc(userID).Collection("posts").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return err
		}
		if _, err := doc.Ref.Delete(ctx); err != nil {
			return err
		}
	}
	return nil
}
