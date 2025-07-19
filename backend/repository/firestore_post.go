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

// FirestorePostRepository は Firestore を使った投稿リポジトリ
type FirestorePostRepository struct {
	Client *firestore.Client // Firestore クライアント
}

// GetAll は posts コレクションの全ドキュメントを取得し、Postスライスで返す
func (r *FirestorePostRepository) GetAll(ctx context.Context) ([]model.Post, error) {
	iter := r.Client.Collection("posts").Documents(ctx) // ドキュメント取得用イテレータ
	defer iter.Stop()                                   // イテレータ解放

	var posts []model.Post
	for {
		doc, err := iter.Next() // 次のドキュメントを取得
		if err != nil {
			break // ドキュメントが無くなったらループ終了
		}

		var post model.Post
		if err := doc.DataTo(&post); err != nil {
			continue // デコード失敗したらスキップ
		}

		post.ID = doc.Ref.ID // FirestoreのドキュメントIDをセット
		posts = append(posts, post)
	}
	return posts, nil
}

// Create は posts コレクションに新しい投稿を追加
func (r *FirestorePostRepository) Create(ctx context.Context, post *model.Post) error {
	docRef, _, err := r.Client.Collection("posts").Add(ctx, post) // 新規ドキュメント追加
	if err != nil {
		return err
	}
	post.ID = docRef.ID // 追加されたドキュメントIDをモデルにセット
	return nil
}

// GetByID は指定IDの投稿を取得
func (r *FirestorePostRepository) GetByID(ctx context.Context, id string) (*model.Post, error) {
	doc, err := r.Client.Collection("posts").Doc(id).Get(ctx) // ドキュメント取得
	if err != nil {
		if status.Code(err) == codes.NotFound { // ドキュメントが無い場合の判定
			return nil, fmt.Errorf("not found")
		}
		return nil, err
	}

	var post model.Post
	if err := doc.DataTo(&post); err != nil { // ドキュメントのデータをモデルに変換
		return nil, err
	}
	post.ID = doc.Ref.ID // ドキュメントIDをセット
	return &post, nil
}

// Update は指定IDの投稿を更新する
func (r *FirestorePostRepository) Update(ctx context.Context, id string, post model.Post) error {
	_, err := r.Client.Collection("posts").Doc(id).Set(ctx, map[string]interface{}{
		"user_id":    post.UserID,
		"title":      post.Title,
		"content":    post.Content,
		"updated_at": time.Now(),
	}, firestore.MergeAll) // MergeAllでフィールドをマージ（上書き）
	return err
}

// Delete は指定IDの投稿を削除する
func (r *FirestorePostRepository) Delete(ctx context.Context, id string) error {
	_, err := r.Client.Collection("posts").Doc(id).Delete(ctx) // ドキュメント削除
	return err
}

// DeleteAll は posts コレクション内の全投稿を一括削除する
func (r *FirestorePostRepository) DeleteAll(ctx context.Context) error {
	iter := r.Client.Collection("posts").Documents(ctx) // ドキュメント取得イテレータ
	defer iter.Stop()                                    // イテレータ解放

	for {
		doc, err := iter.Next() // 次のドキュメント取得
		if err != nil {
			if err == iterator.Done { // ドキュメントがなくなったら終了
				break
			}
			return err // 取得中エラー発生時はそのまま返す
		}
		if _, err := doc.Ref.Delete(ctx); err != nil { // ドキュメント削除
			return err
		}
	}
	return nil
}