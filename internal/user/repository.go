package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, user *User) (uuid.UUID, error)

	ReadById(ctx context.Context, userId uuid.UUID) (*User, error)
	ReadByEmail(ctx context.Context, email string) (*User, error)

	Login(ctx context.Context, email string, password []byte) (*User, error)

	UpdateUserById(ctx context.Context, newUser *User) error

	DeleteUserById(ctx context.Context, userId uuid.UUID) error
}

type Repo struct {
	db *pgxpool.Conn
}

func NewRepository(db *pgxpool.Conn) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(ctx context.Context, user *User) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (r *Repo) ReadById(ctx context.Context, userId uuid.UUID) (*User, error) {
	return nil, nil
}

func (r *Repo) ReadByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}

func (r *Repo) Login(ctx context.Context, email string, password []byte) (*User, error) {
	return nil, nil
}

func (r *Repo) UpdateUserById(ctx context.Context, newUser *User) error {
	return nil
}

func (r *Repo) DeleteUserById(ctx context.Context, userId uuid.UUID) error {
	return nil
}
