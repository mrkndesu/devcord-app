package repository

import (
	"context"

	"devcord-app/model"
)

type PostRepository interface {
	GetAll(ctx context.Context) ([]model.Post, error)
	Create(ctx context.Context, post model.Post) error
}