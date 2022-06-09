package service

import (
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/repository"
)

type Environment struct {
}

func (Environment) FindList(query model.EnvironmentQuery) []model.EnvironmentPO {
	return repository.Environment{}.SelectListBy(query)
}

func (Environment) Create(po model.EnvironmentPO) {
	repository.Environment{}.InsertOne(po)
}
