package repository

import (
	"github.com/czyhome/go-pulsar-manager/model"
)

type Environment struct {
}

func (Environment) SelectListBy(query model.EnvironmentQuery) []model.EnvironmentPO {
	d := dbClient.Model(&model.EnvironmentPO{})
	if len(query.Name) > 0 {
		d.Where(" name like ? ", query.Name+"%")
	}
	var list []model.EnvironmentPO
	d.Find(&list)
	return list
}

func (Environment) InsertOne(po model.EnvironmentPO) {
	dbClient.Create(&po)
}
