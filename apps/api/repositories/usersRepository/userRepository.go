package usersRepository

import (
	"dragon-threads/apps/api/database"
	"dragon-threads/apps/api/entities"
)

func GetUser(query interface{}, args ...interface{}) (*entities.User, error) {
	user := new(entities.User)
	err := database.SQL.Where(query, args).First(user).Error
	return user, err
}

func GetUserList(query string, page int, pageSize int, args []interface{}) ([]entities.User, int64, error) {
	var total int64
	err := database.SQL.Model(&entities.User{}).Where(query, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	users := []entities.User{}
	err = database.SQL.Where(query, args...).Offset(offset).Limit(pageSize).Find(&users).Error
	return users, total, err
}
func CreateUser(user *entities.User) (*entities.User, error) {
	err := database.SQL.Create(user).Error
	return user, err
}

func UpdateUser(user *entities.User) (*entities.User, error) {
	err := database.SQL.Save(user).Error
	return user, err
}

func DeleteUser(user *entities.User) (*entities.User, error) {
	err := database.SQL.Delete(&user).Error
	return user, err
}

func HardDeleteUser(user *entities.User) (*entities.User, error) {
	err := database.SQL.Unscoped().Delete(&user).Error
	return user, err
}
