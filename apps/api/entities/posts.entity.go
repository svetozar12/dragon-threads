package entities

type Post struct {
	Model
	Title       string `json:"title"`
	Content     string `json:"content"`
	UserID      int32  `json:"user_id"`     // Foreign key to User
	SubDragonId int32  `json:"subDragonId"` // Foreign key to SubDragon
	Votes       int32  `json:"votes"`
}
