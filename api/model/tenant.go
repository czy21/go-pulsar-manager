package model

type TenantQuery struct {
	BaseQuery[int64]
	Url string `json:"url"`
}

type TenantDTO struct {
	Name       string `json:"name"`
	Namespaces int    `json:"namespaces"`
}
