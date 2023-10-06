package entities

type PostComment struct {
	Model
	Comment       string `json:"comment"`
	ParentComment int32  `json:"parentComment"` // Reference to other comments
	UserID        int32  `json:"user_id"`       // Foreign key to User
	PostId        int32  `json:"postId"`        // Foreign key to Post
}
