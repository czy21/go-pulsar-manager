package service

import (
	"github.com/czyhome/go-pulsar-manager/constant"
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/go-resty/resty/v2"
)

type Tenant struct {
}

func (Tenant) FindList(query model.TenantQuery) []model.TenantDTO {
	client := resty.New()
	client.SetBaseURL(query.Url)
	req := client.R()
	res, _ := req.Get(constant.PulsarRestPath + "/tenants")
	var tenants []string
	_ = client.JSONUnmarshal(res.Body(), &tenants)
	var tenantDTOs []model.TenantDTO
	for _, t := range tenants {
		var namespaces []string
		res, _ := req.Get(constant.PulsarRestPath + "/namespaces/" + t)
		_ = client.JSONUnmarshal(res.Body(), &namespaces)
		tenantDTOs = append(tenantDTOs,
			model.TenantDTO{
				Name:       t,
				Namespaces: len(namespaces),
			})
	}
	return tenantDTOs
}
