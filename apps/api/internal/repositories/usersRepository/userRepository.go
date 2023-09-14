package usersRepository

import (
	"dragon-threads/apps/api/internal/database"
	"dragon-threads/apps/api/internal/entities"
)

func GetUser(query interface{}, args ...interface{}) (*entities.User, error) {
	user := new(entities.User)
	err := database.SQL.Where(query, args).First(user).Error
	return user, err
}

func GetUserList(userIds []string, args ...interface{}) ([]entities.User, error) {
	products := []entities.User{}
	err := database.SQL.Where("id in (?)", userIds, args).Find(&products).Error
	return products, err
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
