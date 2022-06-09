package model

type TenantQuery struct {
	BaseQuery[int64]
}

type TenantDTO struct {
	Name       string `json:"name"`
	Namespaces int    `json:"namespaces"`
}
