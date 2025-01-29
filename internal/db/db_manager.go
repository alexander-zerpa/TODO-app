package db

import (
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

func (sqldb *SQLiteDB) Init() (err error) {
	sqldb.db, err = gorm.Open(sqlite.Open(sqldb.config.path), &gorm.Config{})
	if err != nil {
		return
	}

	err = sqldb.db.AutoMigrate(&models.Todo{})
	return
}

func (sqldb *SQLiteDB) Add(newTodo *models.Todo) (err error) {
	data, err := sqldb.get(newTodo.ID)
	if err != nil {
		return
	} else if len(data) != 0 {
		return gorm.ErrDuplicatedKey
	}
	err = sqldb.db.Create(newTodo).Error
	return
}

func (sqldb *SQLiteDB) Update(newTodo *models.Todo) (err error) {
	return
}

func (sqldb *SQLiteDB) List(done bool) (data []models.Todo, err error) {
	err = sqldb.db.Find(&data, "done = ?", done).Error
	return
}

func (sqldb *SQLiteDB) ListAll() (data []models.Todo, err error) {
	err = sqldb.db.Find(&data).Error
	return
}

func (sqldb *SQLiteDB) get(id string) (data []models.Todo, err error) {
	err = sqldb.db.Find(&data, "ID = ?", id).Error
	return
}
