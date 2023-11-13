package db

import (
	"ms-user-api/exceptions"
	"ms-user-api/user/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB 
}

func NewDatabase() (entities.Storage, error) {
	db, err := gorm.Open(postgres.Open("user=postgres password=admin dbname=botbank port=5432 sslmode=disable"))
	return &Database{db: db}, err
}

func (orm *Database) Get(userId string) (user *entities.User, err error) {
	if err = orm.db.First(&user, "id = ?", userId).Error; err != nil {
		return nil, err
	}

	return user, err
}

func (orm *Database) Post(req *entities.User) (user *entities.User, err error) {
	if err = orm.db.FirstOrCreate(&user, "id = ?", req.UserId).Error; err != nil {
		return nil, exceptions.New(exceptions.InternalServerError, "internal server error")
	}

	return user, nil
}

func (orm *Database) AutoMigrateSetup() {
	orm.db.AutoMigrate(&entities.User{})
}