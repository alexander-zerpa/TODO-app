package db

import (
	"errors"
	"reflect"
	"testing"
	"todo/internal/models"

	"gorm.io/gorm"
)

func initTest(t *testing.T) (db *SQLiteDB) {
	t.Helper()

	//db = &SQLiteDB{config: DBConfig{path: "db_test.db"}}
	db = &SQLiteDB{config: DBConfig{path: ":memory:"}}
	err := db.Init()
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

	err := db.Add(&expectedTodo)
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

	err = db.Add(&expectedTodo)
	if !errors.Is(err, gorm.ErrDuplicatedKey) {
		t.Errorf("Expected error adding duplicate todo, recived: %v", err)
	}
}

func TestList(t *testing.T) {
	cases := []struct {
		Name      string
		InitialDB []models.Todo
		Input     bool
		Expected  []models.Todo
	}{
		{
			"get done items",
			[]models.Todo{
				{ID: "1", Title: "Test", Description: "Test", Done: false},
				{ID: "2", Title: "Test", Description: "Test", Done: true},
			},
			true,
			[]models.Todo{
				{ID: "2", Title: "Test", Description: "Test", Done: true},
			},
		},
		{
			"get pending items",
			[]models.Todo{
				{ID: "1", Title: "Test", Description: "Test", Done: false},
				{ID: "2", Title: "Test", Description: "Test", Done: true},
			},
			false,
			[]models.Todo{
				{ID: "1", Title: "Test", Description: "Test", Done: false},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			db := populateDB(t, test.InitialDB)
			got, err := db.List(test.Input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("Expected %v, got %v", test.Expected, got)
			}
		})
	}
}

func TestListAll(t *testing.T) {
	cases := []struct {
		Name     string
		Expected []models.Todo
	}{
		{
			"empty List",
			[]models.Todo{},
		},
		{
			"with data",
			[]models.Todo{
				{ID: "1", Title: "Test", Description: "Test", Done: false},
				{ID: "2", Title: "Test", Description: "Test", Done: true},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			db := populateDB(t, test.Expected)
			got, err := db.ListAll()
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("Expected %v, got %v", test.Expected, got)
			}
		})
	}
}

func populateDB(t *testing.T, data []models.Todo) (db *SQLiteDB) {
	db = initTest(t)
	for _, todo := range data {
		err := db.db.Create(&todo).Error
		if err != nil {
			t.Fatalf("Error adding todo: %v", err)
		}
	}
	return
}

func TestUpdate(t *testing.T) {
	db := initTest(t)
	initialTodo := models.Todo{ID: "1", Title: "Test", Description: "Test", Done: false}
	expectedTodo := models.Todo{ID: "1", Title: "Test", Description: "Test", Done: true}

	err := db.db.Create(&initialTodo).Error
	if err != nil {
		t.Errorf("Error adding todo: %v", err)
	}

	err = db.Update(&expectedTodo)
	if err != nil {
		t.Errorf("Error updating todo: %v", err)
	}

	var res models.Todo

	err = db.db.First(&res, "ID = ?", expectedTodo.ID).Error
	if err != nil {
		t.Fatalf("Error getting todo: %v", err)
	}
	if !reflect.DeepEqual(res, expectedTodo) {
		t.Errorf("Expected %v, got %v", expectedTodo, res)
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
