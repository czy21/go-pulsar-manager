package repository

import (
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/web"
)

type Cluster struct {
}

func (Cluster) SelectPage(query model.ClusterQuery) ([]model.ClusterPO, web.PageModel) {
	d := dbClient.Model(&model.ClusterPO{})
	if query.Name != "" {
		d.Where(" name = ? ", query.Name)
	}
	page := web.PageModel{PageIndex: query.PageIndex, PageSize: query.PageSize}
	var list []model.ClusterPO
	d.Count(&page.Total)
	d.Offset((query.PageIndex - 1) * query.PageSize).Limit(query.PageSize).Find(&list)
	return list, page
}

func (Cluster) InsertOne(po model.ClusterPO) {
	dbClient.Create(&po)
}
