package service

import (
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/repository"
)

type Cluster struct {
}

func (Cluster) Paging(query model.ClusterQuery) model.PageResult[model.ClusterPO] {
	list, page := repository.Cluster{}.SelectPage(query)
	return model.PageResult[model.ClusterPO]{Page: page, List: list}
}

func (Cluster) Create(po model.ClusterPO) {
	repository.Cluster{}.InsertOne(po)
}
