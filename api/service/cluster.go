package service

import (
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/repository"
)

type Cluster struct {
}

func (c Cluster) Create(po model.ClusterPO) {
	repository.Cluster{}.InsertOne(po)
}
