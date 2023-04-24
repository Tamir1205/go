package items

import (
	"context"
	"fmt"
	"github.com/Tamir1205/midterm1/internal/storage"
	"strings"
)

type Item struct {
	Id          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Rating      []float64 `db:"rating"`
	CreatedAt   string    `db:"created_at"`
	UpdatedAt   string    `db:"updated_at"`
}

type repository struct {
	db        storage.Storage
	tableName string
}

type Repository interface {
	CreateItem(ctx context.Context, item Item) (int64, error)
	FindItemsByName(ctx context.Context, filter []Filter) ([]Item, error)
	FindItemById(ctx context.Context, id int64) (Item, error)
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

func (r *repository) FindItemsByName(ctx context.Context, filter []Filter) ([]Item, error) {
	query := `SELECT i.id, name, price, description, created_at, updated_at
			  FROM (SELECT i.*, coalesce(avg(r.rating)::numeric(10, 1), 0) as rating
					  FROM item i
							   LEFT JOIN rating r on i.id = r.item_id
					  GROUP BY i.id) as i
			  WHERE 1 = 1
`

	filterQueries, args := r.FilterItem(filter)
	query += filterQueries

	items := make([]Item, 0)
	err := r.db.SelectContext(ctx, &items, query, args...)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *repository) FindItemById(ctx context.Context, id int64) (Item, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, r.tableName)

	var item Item
	err := r.db.GetContext(ctx, &item, query, id)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

type Filter struct {
	Key   string
	Value string
}

func (r *repository) FilterItem(filter []Filter) (string, []interface{}) {
	var arg []interface{}
	query := ""
	for _, v := range filter {
		if v.Key == "name" {
			query += fmt.Sprintf("and lower(name) like lower($%d)\n", len(arg)+1)
			arg = append(arg, "%"+v.Value+"%")
		}
		if v.Key == "price" {
			split := strings.Split(v.Value, "-")
			if len(split) == 2 {
				query += fmt.Sprintf("and price between $%d and $%d\n", len(arg)+1, len(arg)+2)
				arg = append(arg, split[0])
				arg = append(arg, split[1])
			}
		}
		if v.Key == "rating" {
			split := strings.Split(v.Value, "-")
			if len(split) == 2 {
				query += fmt.Sprintf("and rating between $%d and $%d\n", len(arg)+1, len(arg)+2)
				arg = append(arg, split[0])
				arg = append(arg, split[1])
			}
		}
	}
	return query, arg
}
