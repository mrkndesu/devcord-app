package repository

import (
	"context"

	"devcord-app/model"

	"cloud.google.com/go/firestore"
)

type FirestorePostRepository struct {
	Client *firestore.Client
}

func (r *FirestorePostRepository) GetAll(ctx context.Context) ([]model.Post, error) {
	iter := r.Client.Collection("posts").Documents(ctx)
	defer iter.Stop()

	var posts []model.Post
	for {
		doc, err := iter.Next()
		if err != nil {
			break // or check for iterator.Done
		}
		var post model.Post
		if err := doc.DataTo(&post); err != nil {
			continue
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *FirestorePostRepository) Create(ctx context.Context, post model.Post) error {
	_, _, err := r.Client.Collection("posts").Add(ctx, post)
	return err
}