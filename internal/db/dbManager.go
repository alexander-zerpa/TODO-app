package db

import (
	"errors"
	"todo/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDB struct {
	db     *gorm.DB
	config DBConfig
}
type DBConfig struct {
	path string
}

func (sqldb *SQLiteDB) init() (err error) {
	sqldb.db, err = gorm.Open(sqlite.Open(sqldb.config.path), &gorm.Config{})
	if err != nil {
		return
	}

	err = sqldb.db.AutoMigrate(&models.Todo{})
	return
}

func (sqldb *SQLiteDB) add(newTodo *models.Todo) (err error) {
	_, err = sqldb.get(newTodo.ID)
	if err == nil {
		return gorm.ErrDuplicatedKey
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	err = sqldb.db.Create(newTodo).Error
	return
}

func (sqldb *SQLiteDB) get(id string) (data []models.Todo, err error) {
	err = sqldb.db.Find(&data, "ID = ?", id).Error
	return
}

//func (sqldb *SQLiteDB) list(done bool) (data []models.Todo, err error) {
//	err = sqldb.db.Find(&data, "done = ", done).Error
//	return
//}
//
//func (sqldb *SQLiteDB) listAll() (data []models.Todo, err error) {
//	err = sqldb.db.Find(&data).Error
//	return
//}
