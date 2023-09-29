package postsRepository

import (
	"dragon-threads/apps/api/database"
	"dragon-threads/apps/api/entities"
)

func GetPost(query interface{}, args ...interface{}) (*entities.Post, error) {
	post := new(entities.Post)
	err := database.SQL.Where(query, args).First(post).Error
	return post, err
}

func GetPostList(query string, page int, pageSize int, args []interface{}) ([]entities.Post, int64, error) {
	var total int64
	err := database.SQL.Model(&entities.Post{}).Where(query, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	posts := []entities.Post{}
	err = database.SQL.Where(query, args...).Offset(offset).Limit(pageSize).Find(&posts).Error
	return posts, total, err
}

func SearchPostByText(query string, page int, pageSize int) ([]entities.Post, int64, error) {
	var total int64

	likeQuery := "title LIKE ?"
	args := []interface{}{"%" + query + "%"}

	err := database.SQL.Model(&entities.Post{}).Where(likeQuery, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	posts := []entities.Post{}
	err = database.SQL.Where(likeQuery, args...).Offset(offset).Limit(pageSize).Find(&posts).Error
	return posts, total, err
}

func CreatePost(post *entities.Post) (*entities.Post, error) {
	err := database.SQL.Create(post).Error
	return post, err
}

func UpdatePost(post *entities.Post) (*entities.Post, error) {
	err := database.SQL.Save(post).Error
	return post, err
}

func DeletePost(post *entities.Post) (*entities.Post, error) {
	err := database.SQL.Delete(&post).Error
	return post, err
}

func HardDeletePost(post *entities.Post) (*entities.Post, error) {
	err := database.SQL.Unscoped().Delete(&post).Error
	return post, err
}
