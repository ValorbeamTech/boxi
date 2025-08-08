package service

import (
	"context"
	"errors"
	"time"

	"github.com/francis/projectx-api/internal/model"
	"github.com/francis/projectx-api/internal/repository"
	"github.com/francis/projectx-api/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
    userRepo  repository.UserRepository
    jwtSecret string
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string) *AuthService {
    return &AuthService{
        userRepo:  userRepo,
        jwtSecret: jwtSecret,
    }
}

func (s *AuthService) Register(ctx context.Context, req *model.CreateUserRequest) (*model.User, error) {
    // Check if user exists
    existingUser, _ := s.userRepo.GetByEmail(ctx, req.Email)
    if existingUser != nil {
        return nil, errors.New("user already exists")
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    user := &model.User{
        Email:     req.Email,
        FirstName: req.FirstName,
        LastName:  req.LastName,
        Password:  hashedPassword,
    }

    if err := s.userRepo.Create(ctx, user); err != nil {
        return nil, err
    }

    return user, nil
}

func (s *AuthService) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
    user, err := s.userRepo.GetByEmail(ctx, req.Email)
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    if !utils.CheckPasswordHash(req.Password, user.Password) {
        return nil, errors.New("invalid credentials")
    }

    token, err := s.generateToken(user.ID, user.Email)
    if err != nil {
        return nil, err
    }

    refreshToken, err := s.generateRefreshToken(user.ID)
    if err != nil {
        return nil, err
    }

    return &model.LoginResponse{
        Token:        token,
        RefreshToken: refreshToken,
        User:         *user,
    }, nil
}

func (s *AuthService) generateToken(userID int, email string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "email":   email,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
        "iat":     time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) generateRefreshToken(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
        "iat":     time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid token")
        }
        return []byte(s.jwtSecret), nil
    })
}