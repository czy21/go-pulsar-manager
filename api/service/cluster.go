package service

import (
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/repository"
)

type Cluster struct {
}

func (Cluster) FindList(query model.ClusterQuery) []model.ClusterPO {
	return repository.Cluster{}.SelectListBy(query)
}

func (Cluster) Create(po model.ClusterPO) {
	repository.Cluster{}.InsertOne(po)
}
