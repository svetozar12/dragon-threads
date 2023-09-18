package entities

type SubDragon struct {
	Model
	Name            string   `json:"title"`
	Description     string   `json:"description"`
	CommunityTopics []string `json:"communityTopics"`
	UserIds         []int32  `json:"userIds"` // Foreign key to Users
	PostIds         []int32  `json:"postIds"` // Foreign key to Posts
	// Add more fields as needed
}
