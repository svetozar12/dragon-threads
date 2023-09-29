package subDragonRepository

import (
	"dragon-threads/apps/api/database"
	"dragon-threads/apps/api/entities"
)

func GetSubDragon(query interface{}, args ...interface{}) (*entities.SubDragon, error) {
	post := new(entities.SubDragon)
	err := database.SQL.Where(query, args).First(post).Error
	return post, err
}

func GetSubDragonList(query string, page int, pageSize int, args []interface{}) ([]entities.SubDragon, int64, error) {
	var total int64
	err := database.SQL.Model(&entities.SubDragon{}).Where(query, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	subDragon := []entities.SubDragon{}
	err = database.SQL.Where(query, args...).Offset(offset).Limit(pageSize).Find(&subDragon).Error
	return subDragon, total, err
}

func SearchSubDragonByText(query string, page int, pageSize int) ([]entities.SubDragon, int64, error) {
	var total int64

	likeQuery := "name LIKE ?"
	args := []interface{}{"%" + query + "%"}

	err := database.SQL.Model(&entities.SubDragon{}).Where(likeQuery, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	subDragons := []entities.SubDragon{}
	err = database.SQL.Where(likeQuery, args...).Offset(offset).Limit(pageSize).Find(&subDragons).Error
	return subDragons, total, err
}

func CreateSubDragon(post *entities.SubDragon) (*entities.SubDragon, error) {
	err := database.SQL.Create(post).Error
	return post, err
}

func UpdateSubDragon(post *entities.SubDragon) (*entities.SubDragon, error) {
	err := database.SQL.Save(post).Error
	return post, err
}

func DeleteSubDragon(post *entities.SubDragon) (*entities.SubDragon, error) {
	err := database.SQL.Delete(&post).Error
	return post, err
}

func HardDeleteSubDragon(post *entities.SubDragon) (*entities.SubDragon, error) {
	err := database.SQL.Unscoped().Delete(&post).Error
	return post, err
}
