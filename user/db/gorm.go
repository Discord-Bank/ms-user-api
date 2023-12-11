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

func NewDatabase(dsn string) (entities.Storage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return &Database{db: db}, err
}

func (orm *Database) List(userIds []string, limit int, page int) ([]entities.User, error) {
	var user = []entities.User{}
	var err error
	if len(userIds) < 1 {
		if page == 0 {
			page = 1
		}
		offset := (page - 1) * limit
		err = orm.db.Offset(offset).Limit(limit).Find(&user).Error
	} else {
		err = orm.db.Where("id IN ?", userIds).Find(&user).Error
	}

	return user, err
}

func (orm *Database) Post(req *entities.User) (user *entities.User, err error) {
	if err := orm.db.Debug().Create(&req).Error; err != nil {
		return user, exceptions.New(exceptions.InternalServerError, "internal server error")
	}

	return req, nil
}

func (orm *Database) Patch(req *entities.User) error {
	if err := orm.db.Save(&req).Error; err != nil {
		return exceptions.New(exceptions.NotFound, "user with id "+req.Id+" not found")
	}
	return nil
}

func (orm *Database) Delete(req *entities.User) error {
	if err := orm.db.Delete(&req).Error; err != nil {
		return exceptions.New(exceptions.NotFound, "user with id "+req.Id+" not found")
	}
	return nil
}

func (orm *Database) AutoMigrateSetup() {
	orm.db.AutoMigrate(&entities.User{})
}
