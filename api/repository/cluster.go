package repository

import (
	"github.com/czyhome/go-pulsar-manager/model"
)

type Cluster struct {
}

func (Cluster) SelectPage(query model.ClusterQuery) ([]model.ClusterPO, model.PageModel) {
	d := dbClient.Model(&model.ClusterPO{})
	if len(query.Name) > 0 {
		d.Where(" name like ? ", query.Name+"%")
	}
	return Base[model.ClusterPO]{DB: d}.SelectPage(query.PageIndex, query.PageSize)
}

func (Cluster) InsertOne(po model.ClusterPO) {
	dbClient.Create(&po)
}
