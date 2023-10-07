package postVoteRepository

import (
	"dragon-threads/apps/api/database"
	"dragon-threads/apps/api/entities"
)

func GetPostVote(query interface{}, args ...interface{}) (*entities.PostVote, error) {
	post := new(entities.PostVote)
	err := database.SQL.Where(query, args).First(post).Error
	return post, err
}

func GetPostVoteList(query string, page int, pageSize int, args []interface{}) ([]entities.PostVote, int64, error) {
	var total int64
	err := database.SQL.Model(&entities.PostVote{}).Where(query, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	postVotes := []entities.PostVote{}
	err = database.SQL.Where(query, args...).Offset(offset).Limit(pageSize).Find(&postVotes).Error
	return postVotes, total, err
}

func CreatePostVote(post *entities.PostVote) (*entities.PostVote, error) {
	err := database.SQL.Create(post).Error
	return post, err
}

func UpdatePostVote(post *entities.PostVote) (*entities.PostVote, error) {
	err := database.SQL.Save(post).Error
	return post, err
}

func DeletePostVote(post *entities.PostVote) (*entities.PostVote, error) {
	err := database.SQL.Delete(&post).Error
	return post, err
}

func HardDeletePostVote(post *entities.PostVote) (*entities.PostVote, error) {
	err := database.SQL.Unscoped().Delete(&post).Error
	return post, err
}
