package repository

import "newgens/models"

type MT202Repo interface {
	GetData() ([]*models.MT202, error)
	InsertData(data *models.MT202) error
}
