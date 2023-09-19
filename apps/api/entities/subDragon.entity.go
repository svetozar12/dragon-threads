package entities

type SubDragon struct {
	Model
	Name            string                     `json:"name"`
	Description     string                     `json:"description"`
	CommunityTopics []SubDragonCommunityTopics `json:"communityTopics"`
	// Add more fields as needed
}

type SubDragonCommunityTopics struct {
	Model
	Name        string `json:"name"`
	SubDragonId uint   `json:"subDragonId"` // Foreign key to SubDragon
}
