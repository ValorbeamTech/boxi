package model

import (
    "time"
)

type User struct {
    ID        int       `json:"id" db:"id"`
    Email     string    `json:"email" db:"email" validate:"required,email"`
    FirstName string    `json:"first_name" db:"first_name" validate:"required"`
    LastName  string    `json:"last_name" db:"last_name" validate:"required"`
    Password  string    `json:"-" db:"password_hash"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateUserRequest struct {
    Email     string `json:"email" validate:"required,email"`
    FirstName string `json:"first_name" validate:"required"`
    LastName  string `json:"last_name" validate:"required"`
    Password  string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
    Token        string `json:"token"`
    RefreshToken string `json:"refresh_token"`
    User         User   `json:"user"`
}