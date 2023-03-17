package users

import (
	"context"
	"fmt"
	"github.com/Tamir1205/midterm1/internal/storage"
)

type User struct {
	ID        int64  `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type repository struct {
	db        storage.Storage
	tableName string
}

type Repository interface {
	CreateUser(ctx context.Context, user User) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

func NewRepository(db storage.Storage) Repository {
	return &repository{
		db:        db,
		tableName: "users",
	}
}

func (r *repository) CreateUser(ctx context.Context, user User) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (email, password, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id", r.tableName)

	var id int64
	err := r.db.QueryRowxContext(ctx, query, user.Email, user.Password, user.FirstName, user.LastName).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) GetUser(ctx context.Context, id int64) (User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", r.tableName)

	var user User
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", r.tableName)

	var user User
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
