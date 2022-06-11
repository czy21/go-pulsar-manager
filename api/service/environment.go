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

func (Environment) SimpleItems() []model.SimpleItemModel[int64] {
	l := repository.Environment{}.SelectListBy(model.EnvironmentQuery{})
	var items []model.SimpleItemModel[int64]
	for _, t := range l {
		items = append(items, model.SimpleItemModel[int64]{Label: t.Name, Value: t.Id, Extra: map[string]interface{}{"serviceUrl": t.Url}})
	}
	return items
}

func (Environment) Create(po model.EnvironmentPO) {
	repository.Environment{}.InsertOne(po)
}
