package model

type Post struct {
	Title   string `json:"title" firestore:"title"`
	Content string `json:"content" firestore:"content"`
}