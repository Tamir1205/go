package rating

import (
	"context"
	"fmt"
	"github.com/Tamir1205/midterm1/internal/storage"
)

type Rating struct {
	ID        int64  `db:"id"`
	UserId    int64  `db:"user_id"`
	ItemId    int64  `db:"item_id"`
	Rating    uint   `db:"rating"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type repository struct {
	db        storage.Storage
	tableName string
}

type Repository interface {
	Create(ctx context.Context, rating Rating) (int64, error)
	GetRatingsByItemId(ctx context.Context, itemId int64) ([]Rating, error)
}

func NewRepository(db storage.Storage) Repository {
	return &repository{
		db:        db,
		tableName: "rating",
	}
}

func (r *repository) Create(ctx context.Context, rating Rating) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, item_id, rating) VALUES ($1, $2, $3) RETURNING id", r.tableName)

	var id int64
	err := r.db.QueryRowxContext(ctx, query, rating.UserId, rating.ItemId, rating.Rating).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) GetRatingsByItemId(ctx context.Context, itemId int64) ([]Rating, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE item_id = $1`, r.tableName)

	var comment []Rating
	err := r.db.SelectContext(ctx, &comment, query, itemId)
	if err != nil {
		return comment, err
	}

	return comment, nil
}
