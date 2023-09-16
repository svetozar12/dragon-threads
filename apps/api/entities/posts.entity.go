package entities

type Post struct {
	Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"` // Foreign key to User
	// Add more fields as needed
}
