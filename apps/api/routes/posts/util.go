package posts

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/repositories/postsRepository"
)

func getPostsByQuery(page int, pageSize int, subDragonId int) ([]entities.Post, int64, error) {
	// Retrieve user list based on query parameters
	postList, total, err := postsRepository.GetPostList("sub_dragon_id=?", page, pageSize, []interface{}{subDragonId})
	return postList, total, err
}

func updatePostFields(existingPost *entities.Post, updatedPost UpdatePostSchema) {
	if updatedPost.Title != "" {
		existingPost.Title = updatedPost.Title
	}

	if updatedPost.Content != "" {
		existingPost.Content = updatedPost.Content
	}
}
