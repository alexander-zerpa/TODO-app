package db

import (
	"todo/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDB struct {
	db     *gorm.DB
	config DBConfig
}

type DBConfig struct {
	Path string
}

func NewSQLiteDB(config DBConfig) (db *sqliteDB) {
	db = &sqliteDB{config: config}
	db.init()
	return
}

func (sqldb *sqliteDB) init() (err error) {
	sqldb.db, err = gorm.Open(sqlite.Open(sqldb.config.Path), &gorm.Config{})
	if err != nil {
		return
	}

	err = sqldb.db.AutoMigrate(&models.Todo{})
	return
}

func (sqldb *sqliteDB) Add(newTodo *models.Todo) (err error) {
	data, err := sqldb.Get(newTodo.ID)
	if err != nil {
		return
	} else if len(data) != 0 {
		return gorm.ErrDuplicatedKey
	}
	err = sqldb.db.Create(newTodo).Error
	return
}

func (sqldb *sqliteDB) Update(newTodo *models.Todo) (err error) {
	oldTodo, err := sqldb.Get(newTodo.ID)
	if err != nil {
		return
	} else if len(oldTodo) != 1 {
		return gorm.ErrInvalidData
	}
	sqldb.db.Model(&oldTodo[0]).Updates(newTodo)
	return
}

func (sqldb *sqliteDB) List(done bool) (data []models.Todo, err error) {
	err = sqldb.db.Find(&data, "done = ?", done).Error
	return
}

func (sqldb *sqliteDB) ListAll() (data []models.Todo, err error) {
	err = sqldb.db.Find(&data).Error
	return
}

func (sqldb *sqliteDB) Get(id string) (data []models.Todo, err error) {
	err = sqldb.db.Find(&data, "ID = ?", id).Error
	return
}
