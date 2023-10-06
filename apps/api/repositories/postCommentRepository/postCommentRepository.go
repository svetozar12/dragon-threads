package postCommentRepository

import (
	"dragon-threads/apps/api/database"
	"dragon-threads/apps/api/entities"
)

func GetPostComment(query interface{}, args ...interface{}) (*entities.PostComment, error) {
	postComment := new(entities.PostComment)
	err := database.SQL.Where(query, args).First(postComment).Error
	return postComment, err
}

func GetPostCommentList(query string, page int, pageSize int, args []interface{}) ([]entities.PostComment, int64, error) {
	var total int64
	err := database.SQL.Model(&entities.PostComment{}).Where(query, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	postComments := []entities.PostComment{}
	err = database.SQL.Where(query, args...).Offset(offset).Limit(pageSize).Find(&postComments).Error
	return postComments, total, err
}

func CreatePostComment(postComment *entities.PostComment) (*entities.PostComment, error) {
	err := database.SQL.Create(postComment).Error
	return postComment, err
}

func UpdatePostComment(postComment *entities.PostComment) (*entities.PostComment, error) {
	err := database.SQL.Save(postComment).Error
	return postComment, err
}

func DeletePostComment(postComment *entities.PostComment) (*entities.PostComment, error) {
	err := database.SQL.Delete(&postComment).Error
	return postComment, err
}

func HardDeletePostComment(postComment *entities.PostComment) (*entities.PostComment, error) {
	err := database.SQL.Unscoped().Delete(&postComment).Error
	return postComment, err
}
