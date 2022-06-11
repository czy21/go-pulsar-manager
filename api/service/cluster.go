package service

import (
	"github.com/czyhome/go-pulsar-manager/constant"
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/util"
)

type Cluster struct {
}

func (Cluster) FindList(query model.ClusterQuery) []model.ClusterDTO {
	if query.ServiceUrl == "" {
		return []model.ClusterDTO{}
	}
	c := util.HttpUtil{}.NewClient()
	c.SetBaseURL(query.ServiceUrl)
	var strings []string
	c.Get(constant.PulsarRestPath+"/clusters", &strings)
	var rets []model.ClusterDTO
	for _, t := range strings {
		var config map[string]interface{}
		c.Get(constant.PulsarRestPath+"/clusters/"+t, &config)
		var brokers []string
		c.Get(constant.PulsarRestPath+"/brokers/"+t, &brokers)
		rets = append(rets,
			model.ClusterDTO{
				Name:    t,
				Brokers: len(brokers),
				Config:  config,
			})
	}
	return rets
}
