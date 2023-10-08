package entities

type PostVote struct {
	Model
	UserID int32 `json:"user_id"` // Foreign key to User
	PostID int32 `json:"post_id"` // Foreign key to Posts
	Vote   int32 `json:"vote"`
}
