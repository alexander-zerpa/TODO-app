package db

import (
    "todo/internal/models"

    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type SQLiteDB struct {
    db *gorm.DB
    config DBConfig
}
type DBConfig struct {
    path string
}

type TodoRecord struct {
    gorm.Model
    todo models.Todo
}

func (sqldb *SQLiteDB) init()  {
    var err error
    sqldb.db, err = gorm.Open(sqlite.Open(sqldb.config.path), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    sqldb.db.AutoMigrate(&TodoRecord{})
}

func (sqldb *SQLiteDB) add(newTodo models.Todo) err error {
    sqldb.db.Create(&TodoRecord{todo: newTodo})
    return
}
