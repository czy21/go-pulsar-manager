package service

import (
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/repository"
	"github.com/czyhome/go-pulsar-manager/web"
)

type Cluster struct {
}

func (Cluster) Paging(query model.ClusterQuery) web.PageResult[model.ClusterPO] {
	list, page := repository.Cluster{}.SelectPage(query)
	return web.PageResult[model.ClusterPO]{Page: page, List: list}
}

func (Cluster) Create(po model.ClusterPO) {
	repository.Cluster{}.InsertOne(po)
}
