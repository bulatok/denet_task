package storage

import "github.com/bulatok/denet_task/internal/models"

type Files interface {
	Get(name string) (*models.File, error)
	Set(file *models.File) error
	All() ([]models.FileOut, error)
}
