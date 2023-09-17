package entities

type Post struct {
	Model
	Title       string `json:"title"`
	Content     string `json:"content"`
	UserID      uint   `json:"user_id"`     // Foreign key to User
	SubDragonId int32  `json:"subDragonId"` // Foreign key to SubDragon
	// Add more fields as needed
}
