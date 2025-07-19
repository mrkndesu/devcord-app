package model

// 後日作成
// User はユーザーデータのモデル構造体
type User struct {
	ID           string `json:"id"`            // ユーザーID
	Name         string `json:"name"`          // ユーザー名
	AvatarURL    string `json:"avatar_url"`    // アバター画像URL
	Description  string `json:"description"`   // ユーザー説明文
	BirthDate    string `json:"birth_date"`    // 生年月日（文字列型）
	CreatedYear  int    `json:"created_year"`  // 登録年
	CreatedMonth int    `json:"created_month"` // 登録月
}