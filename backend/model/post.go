package model

import "time"

// Post は投稿情報のモデル構造体
// Firestore のドキュメントに対応し、JSONとFirestoreのフィールド名を指定
type Post struct {
	ID        string    `json:"id,omitempty" firestore:"-"`        // ドキュメントID（Firestoreには保存せずレスポンスに含む）
	UserID    string    `json:"user_id" firestore:"user_id"`       // 投稿者ユーザーID
	Title     string    `json:"title" firestore:"title"`           // 投稿タイトル
	Content   string    `json:"content" firestore:"content"`       // 投稿内容
	CreatedAt time.Time `json:"created_at" firestore:"created_at"` // 作成日時（ISO 8601形式）
}
