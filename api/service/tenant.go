package service

import (
	"github.com/czyhome/go-pulsar-manager/constant"
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/util"
)

type Tenant struct {
}

func (Tenant) FindList(query model.TenantQuery) []model.TenantDTO {
	if query.ServiceUrl == "" {
		return []model.TenantDTO{}
	}
	c := util.HttpUtil{}.NewClient()
	c.SetBaseURL(query.ServiceUrl)
	var strings []string
	c.Get(constant.PulsarRestPath+"/tenants", &strings)
	var tenantDTOs []model.TenantDTO
	for _, t := range strings {
		var namespaces []string
		c.Get(constant.PulsarRestPath+"/namespaces/"+t, &namespaces)
		tenantDTOs = append(tenantDTOs,
			model.TenantDTO{
				Name:       t,
				Namespaces: len(namespaces),
			})
	}
	return tenantDTOs
}
