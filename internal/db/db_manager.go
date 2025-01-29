package db

import "todo/internal/models"

type TodoDBManager interface {
    Add(*models.Todo) error
    Update(*models.Todo) error
    List(bool) ([]models.Todo, error)
    ListAll() ([]models.Todo, error)
    Get(string) ([]models.Todo, error)
}
