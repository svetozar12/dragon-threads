package database

import (
	"dragon-threads/apps/api/internal/entities"
	"dragon-threads/apps/api/internal/pkg/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SQL *gorm.DB

func Open() {
	var err error
	SQL, err = gorm.Open(postgres.Open(env.Envs.POSTGRES_CONNECTION_STRING), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic(err)
	}
	SQL.AutoMigrate(&entities.User{})
}
