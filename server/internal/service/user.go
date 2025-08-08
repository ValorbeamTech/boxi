package service

import (
	"context"
	"math"

	"github.com/francis/projectx-api/internal/model"
	"github.com/francis/projectx-api/internal/repository"
)

type UserService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) GetByID(ctx context.Context, id int) (*model.User, error) {
    return s.userRepo.GetByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, user *model.User) error {
    return s.userRepo.Update(ctx, user)
}

func (s *UserService) GetUsers(ctx context.Context, page, limit int) (*model.PaginatedResponse, error) {
    offset := (page - 1) * limit
    users, err := s.userRepo.List(ctx, limit, offset)
    if err != nil {
        return nil, err
    }

    total, err := s.userRepo.Count(ctx)
    if err != nil {
        return nil, err
    }

    totalPages := int(math.Ceil(float64(total) / float64(limit)))

    return &model.PaginatedResponse{
        Data:       users,
        Page:       page,
        Limit:      limit,
        Total:      total,
        TotalPages: totalPages,
    }, nil
}
