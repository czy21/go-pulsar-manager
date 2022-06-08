package repository

import (
	"github.com/czyhome/go-pulsar-manager/model"
)

type Cluster struct {
}

func (Cluster) SelectListBy(query model.ClusterQuery) []model.ClusterPO {
	d := dbClient.Model(&model.ClusterPO{})
	if len(query.Name) > 0 {
		d.Where(" name like ? ", query.Name+"%")
	}
	var list []model.ClusterPO
	d.Find(&list)
	return list
}

func (Cluster) InsertOne(po model.ClusterPO) {
	dbClient.Create(&po)
}
