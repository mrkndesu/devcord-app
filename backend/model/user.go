package model

// User はユーザー情報のモデル構造体
// Firestore のドキュメントに対応し、JSONとFirestoreのフィールド名を指定
type User struct {
	ID           string `json:"id" firestore:"-"`                    // ドキュメントID（Firestoreには保存しない）
	Handle       string `json:"handle" firestore:"handle"`           // 表示用ユーザーID（例: @xxxx）
	Name         string `json:"name" firestore:"name"`               // 表示名
	Email        string `json:"email" firestore:"email"`             // メールアドレス
	Password     string `json:"password" firestore:"password"`       // ハッシュ化パスワード
	AvatarURL    string `json:"avatar_url" firestore:"avatar_url"`   // アバター画像URL
	Description  string `json:"description" firestore:"description"` // 自己紹介文
	BirthDate    string `json:"birth_date" firestore:"birth_date"`   // 生年月日（"YYYY-MM-DD"形式）
	CreatedYear  int    `json:"created_year" firestore:"created_year"`   // 登録年（例: 2025）
	CreatedMonth int    `json:"created_month" firestore:"created_month"` // 登録月（例: 7）
}
