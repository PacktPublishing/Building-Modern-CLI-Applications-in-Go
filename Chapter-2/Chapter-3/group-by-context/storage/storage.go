package storage

import (
	"audiofile/models"
)

type Storage interface {
	Upload(bytes []byte, tag string) (int, error)
	List() ([]models.Audio, error)
	Delete(id string, tag string) error
}
