package votes

type PostVoteSchema struct {
	PostID      int32 `json:"post_id" validate:"required"` // Foreign key to User
	UserID      int32 `json:"user_id" validate:"required"` // Foreign key to User
	SubDragonId int32 `json:"subDragonId" validate:"required"`
}
