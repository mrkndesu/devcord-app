package model

import "time"

// Post は投稿データのモデル構造体
type Post struct {
    ID        string    `json:"id,omitempty" firestore:"-"`     // ドキュメントID（Firestoreには保存しない）
    UserID    string    `json:"user_id" firestore:"user_id"`     // ユーザーID
    Title     string    `json:"title" firestore:"title"`         // 投稿タイトル
    Content   string    `json:"content" firestore:"content"`     // 投稿内容
    CreatedAt time.Time `json:"created_at" firestore:"created_at"` // 作成日時
}