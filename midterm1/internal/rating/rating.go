package rating

type CreateRatingDto struct {
	ItemId int64 `json:"item_id"`
	UserId int64 `json:"user_id"`
	Rating uint  `json:"rating"`
}
