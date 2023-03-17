package items

import (
	"context"
	"fmt"
	"github.com/Tamir1205/midterm1/internal/storage"
)

type Item struct {
	Id          int64   `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

type repository struct {
	db        storage.Storage
	tableName string
}

type Repository interface {
	CreateItem(ctx context.Context, item Item) (int64, error)
	FindItemsByName(ctx context.Context, name string) ([]Item, error)
}

func NewRepository(db storage.Storage) Repository {
	return &repository{
		db:        db,
		tableName: "item",
	}
}

func (r *repository) CreateItem(ctx context.Context, item Item) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, description, price) VALUES ($1, $2, $3) RETURNING id", r.tableName)

	var id int64
	err := r.db.QueryRowxContext(ctx, query, item.Name, item.Description, item.Price).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) FindItemsByName(ctx context.Context, name string) ([]Item, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE lower(name) like lower($1)`, r.tableName)

	item := make([]Item, 0)
	err := r.db.SelectContext(ctx, &item, query, fmt.Sprint("%"+name+"%"))
	if err != nil {
		return nil, err
	}

	return item, nil
}
