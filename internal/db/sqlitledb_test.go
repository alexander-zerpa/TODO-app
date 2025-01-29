package db

import (
	"errors"
	"reflect"
	"testing"
	"todo/internal/models"

	"gorm.io/gorm"
)

type TestDB struct {
	InitialDB []models.Todo
	FinalDB   []models.Todo
}

func TestInit(t *testing.T) {
	initTest(t)
}

func TestAdd(t *testing.T) {
	cases := []struct {
		testDB   TestDB
		Name     string
		Input    models.Todo
		Expected error
	}{
		{
			TestDB{
				[]models.Todo{},
				[]models.Todo{
					{ID: "1", Title: "Test", Description: "Test", Done: false},
				},
			},
			"Add element",
			models.Todo{ID: "1", Title: "Test", Description: "Test", Done: false},
			nil,
		},
		{
			TestDB{
				[]models.Todo{
					{ID: "1", Title: "Test", Description: "Test", Done: false},
				},
				[]models.Todo{
					{ID: "1", Title: "Test", Description: "Test", Done: false},
				},
			},
			"Add existent element",
			models.Todo{ID: "1", Title: "Test", Description: "Test", Done: false},
			gorm.ErrDuplicatedKey,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			db := populateDB(t, test.testDB.InitialDB)
			defer validateDB(t, db, test.testDB.FinalDB)

			err := db.Add(&test.Input)
			if !errors.Is(test.Expected, err) {
				t.Errorf("Expected error %v, got %v", test.Expected, err)
			}
		})
	}
}

func TestList(t *testing.T) {
	testDB := TestDB{[]models.Todo{
		{ID: "1", Title: "Test", Description: "Test", Done: false},
		{ID: "2", Title: "Test", Description: "Test", Done: true},
	},
		[]models.Todo{
			{ID: "1", Title: "Test", Description: "Test", Done: false},
			{ID: "2", Title: "Test", Description: "Test", Done: true},
		},
	}

	cases := []struct {
		testDB   TestDB
		Name     string
		Input    bool
		Expected []models.Todo
	}{
		{
			testDB,
			"get done items",
			true,
			[]models.Todo{
				{ID: "2", Title: "Test", Description: "Test", Done: true},
			},
		},
		{
			testDB,
			"get pending items",
			false,
			[]models.Todo{
				{ID: "1", Title: "Test", Description: "Test", Done: false},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			db := populateDB(t, test.testDB.InitialDB)
			defer validateDB(t, db, test.testDB.FinalDB)

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
			defer validateDB(t, db, test.Expected)

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

func TestUpdate(t *testing.T) {
	cases := []struct {
		testDB   TestDB
		Name     string
		Input    models.Todo
		Expected error
	}{
		{
			TestDB{
				[]models.Todo{
					{ID: "1", Title: "Test", Description: "Test", Done: false},
				},
				[]models.Todo{
					{ID: "1", Title: "Test", Description: "Test", Done: true},
				},
			},
			"Update element",
			models.Todo{ID: "1", Title: "Test", Description: "Test", Done: true},
			nil,
		},
		{
			TestDB{
				[]models.Todo{
					{ID: "1", Title: "Test", Description: "Test", Done: false},
				},
				[]models.Todo{
					{ID: "1", Title: "Test", Description: "Test", Done: false},
				},
			},
			"Update nonexistent element",
			models.Todo{ID: "7", Title: "Test", Description: "Test", Done: true},
			gorm.ErrInvalidData,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			db := populateDB(t, test.testDB.InitialDB)
			defer validateDB(t, db, test.testDB.FinalDB)

			err := db.Update(&test.Input)
			if !errors.Is(test.Expected, err) {
				t.Errorf("Expected error %v, got %v", test.Expected, err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		testDB        TestDB
		Name          string
		Input         string
		Expected      []models.Todo
		ExpectedError error
	}{
		{
			TestDB{
				[]models.Todo{
					{ID: "1", Title: "Test", Description: "Test", Done: false},
				},
				[]models.Todo{
					{ID: "1", Title: "Test", Description: "Test", Done: false},
				},
			},
			"Get element",
			"1",
			[]models.Todo{{ID: "1", Title: "Test", Description: "Test", Done: false}},
			nil,
		},
		{
			TestDB{
				[]models.Todo{
					{ID: "1", Title: "Test", Description: "Test", Done: false},
				},
				[]models.Todo{
					{ID: "1", Title: "Test", Description: "Test", Done: false},
				},
			},
			"Get element",
			"7",
			[]models.Todo{},
			nil,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			db := populateDB(t, test.testDB.InitialDB)
			defer validateDB(t, db, test.testDB.FinalDB)

			res, err := db.Get(test.Input)
			if !errors.Is(test.ExpectedError, err) {
				t.Errorf("Expected error %v, got %v", test.Expected, err)
			} else if !reflect.DeepEqual(res, test.Expected) {
				t.Errorf("Expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func initTest(t *testing.T) (db *sqliteDB) {
	t.Helper()

	//db = &SQLiteDB{config: DBConfig{path: "db_test.db"}}
	db = &sqliteDB{config: DBConfig{Path: ":memory:"}}
	err := db.init()
	if err != nil {
		t.Fatalf("Error initializing database: %v", err)
	}
	return
}

func populateDB(t *testing.T, data []models.Todo) (db *sqliteDB) {
	db = initTest(t)
	for _, todo := range data {
		err := db.db.Create(&todo).Error
		if err != nil {
			t.Fatalf("Error adding todo: %v", err)
		}
	}
	return
}

func validateDB(t *testing.T, db *sqliteDB, ExpectedDB []models.Todo) {
	var data []models.Todo
	err := db.db.Find(&data).Error
	if err != nil {
		t.Fatalf("Error validating expectedDB: %v", err)
	} else if !reflect.DeepEqual(data, ExpectedDB) {
		t.Errorf("Expected %v, got %v", ExpectedDB, data)
	}
}
