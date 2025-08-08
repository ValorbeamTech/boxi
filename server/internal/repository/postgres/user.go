package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/francis/projectx-api/internal/model"
	"github.com/francis/projectx-api/internal/repository"
)

type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
    query := `
        INSERT INTO users (email, first_name, last_name, password_hash, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`
    
    now := time.Now()
    user.CreatedAt = now
    user.UpdatedAt = now

    return r.db.QueryRowContext(ctx, query,
        user.Email, user.FirstName, user.LastName, user.Password,
        user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
    user := &model.User{}
    query := `
        SELECT id, email, first_name, last_name, password_hash, created_at, updated_at
        FROM users WHERE id = $1`

    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &user.ID, &user.Email, &user.FirstName, &user.LastName,
        &user.Password, &user.CreatedAt, &user.UpdatedAt)

    if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
    user := &model.User{}
    query := `
        SELECT id, email, first_name, last_name, password_hash, created_at, updated_at
        FROM users WHERE email = $1`

    err := r.db.QueryRowContext(ctx, query, email).Scan(
        &user.ID, &user.Email, &user.FirstName, &user.LastName,
        &user.Password, &user.CreatedAt, &user.UpdatedAt)

    if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
    query := `
        UPDATE users 
        SET email = $2, first_name = $3, last_name = $4, updated_at = $5
        WHERE id = $1`
    
    user.UpdatedAt = time.Now()
    _, err := r.db.ExecContext(ctx, query,
        user.ID, user.Email, user.FirstName, user.LastName, user.UpdatedAt)
    return err
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
    query := `DELETE FROM users WHERE id = $1`
    _, err := r.db.ExecContext(ctx, query, id)
    return err
}

func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*model.User, error) {
    query := `
        SELECT id, email, first_name, last_name, created_at, updated_at
        FROM users ORDER BY created_at DESC LIMIT $1 OFFSET $2`

    rows, err := r.db.QueryContext(ctx, query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*model.User
    for rows.Next() {
        user := &model.User{}
        err := rows.Scan(&user.ID, &user.Email, &user.FirstName,
            &user.LastName, &user.CreatedAt, &user.UpdatedAt)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func (r *userRepository) Count(ctx context.Context) (int64, error) {
    var count int64
    query := `SELECT COUNT(*) FROM users`
    err := r.db.QueryRowContext(ctx, query).Scan(&count)
    return count, err
}