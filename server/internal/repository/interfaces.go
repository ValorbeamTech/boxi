package repository

import (
	"context"

	"github.com/francis/projectx-api/internal/model"
)

type UserRepository interface {
    Create(ctx context.Context, user *model.User) error
    GetByID(ctx context.Context, id int) (*model.User, error)
    GetByEmail(ctx context.Context, email string) (*model.User, error)
    Update(ctx context.Context, user *model.User) error
    Delete(ctx context.Context, id int) error
    List(ctx context.Context, limit, offset int) ([]*model.User, error)
    Count(ctx context.Context) (int64, error)
}