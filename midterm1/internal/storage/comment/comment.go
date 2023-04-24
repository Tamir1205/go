package comment

import (
	"context"
	"fmt"
	"github.com/Tamir1205/midterm1/internal/storage"
)

type Comment struct {
	ID        int64  `db:"id"`
	UserId    int64  `db:"user_id"`
	ItemId    int64  `db:"item_id"`
	Content   string `db:"content"`
	ParentId  *int64 `db:"parent_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type repository struct {
	db        storage.Storage
	tableName string
}

type Repository interface {
	CreateComment(ctx context.Context, comment Comment) (int64, error)
	FindCommentByItemId(ctx context.Context, itemId int64) ([]Comment, error)
}

func NewRepository(db storage.Storage) Repository {
	return &repository{
		db:        db,
		tableName: "comment",
	}
}

func (r *repository) CreateComment(ctx context.Context, comment Comment) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, item_id, content, parent_id) VALUES ($1, $2, $3, $4) RETURNING id", r.tableName)

	var id int64
	err := r.db.QueryRowxContext(ctx, query, comment.UserId, comment.ItemId, comment.Content, comment.ParentId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) FindCommentByItemId(ctx context.Context, id int64) ([]Comment, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE item_id = $1`, r.tableName)

	var comment []Comment
	err := r.db.SelectContext(ctx, &comment, query, id)
	if err != nil {
		return comment, err
	}

	return comment, nil
}
