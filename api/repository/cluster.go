package repository

import "github.com/czyhome/go-pulsar-manager/model"

type Cluster struct {
}

func (Cluster) InsertOne(po model.ClusterPO) {
	dbClient.Create(&po)
}
