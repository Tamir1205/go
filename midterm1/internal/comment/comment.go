package comment

type CreateCommentDto struct {
	ItemId   int64  `json:"item_id"`
	UserId   int64  `json:"user_id"`
	Content  string `json:"content"`
	ParentId *int64 `json:"parent_id"`
}

type Comment struct {
	ID       int64     `json:"id"`
	UserId   int64     `json:"user_id"`
	ItemId   int64     `json:"item_id"`
	Content  string    `json:"content"`
	ParentId *int64    `json:"parent_id"`
	Children []Comment `json:"children"`
}
