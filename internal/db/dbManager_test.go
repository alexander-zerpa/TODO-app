package db

import (
	"reflect"
	"testing"
	"todo/internal/models"
)

func initTest(t *testing.T) (db *SQLiteDB) {
	t.Helper()

	//db = &SQLiteDB{config: DBConfig{path: "db_test.db"}}
	db = &SQLiteDB{config: DBConfig{path: ":memory:"}}
	err := db.init()
	if err != nil {
		t.Fatalf("Error initializing database: %v", err)
	}
	return
}

func TestInit(t *testing.T) {
	initTest(t)
}

func TestAdd(t *testing.T) {
	db := initTest(t)
	expectedTodo := models.Todo{ID: "1", Title: "Test", Description: "Test", Done: false}

	err := db.add(&expectedTodo)
	if err != nil {
		t.Errorf("Error adding todo: %v", err)
	}
	var res models.Todo

	err = db.db.First(&res, "ID = ?", expectedTodo.ID).Error
	if err != nil {
		t.Fatalf("Error getting todo: %v", err)
	}
	if !reflect.DeepEqual(res, expectedTodo) {
		t.Errorf("Expected %v, got %v", expectedTodo, res)
	}

	err = db.add(&expectedTodo)
	if err == nil {
		t.Errorf("Expected error adding duplicate todo")
	}
}

func TestGet(t *testing.T) {
	db := initTest(t)
	expectedTodo := models.Todo{ID: "1", Title: "Test", Description: "Test", Done: false}
	err := db.db.Create(expectedTodo).Error

	if err != nil {
		t.Fatalf("Error adding todo: %v", err)
	}

	res, err := db.get(expectedTodo.ID)
	if err != nil {
		t.Fatalf("Error getting todo: %v", err)
	}
	if !reflect.DeepEqual(res[0], expectedTodo) {
		t.Errorf("Expected %v, got %v", expectedTodo, res)
	}
}
