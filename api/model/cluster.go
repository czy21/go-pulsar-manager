package model

type ClusterDTO struct {
	Name    string                 `json:"name"`
	Brokers int                    `json:"brokers"`
	Config  map[string]interface{} `json:"config"`
}

type ClusterQuery struct {
	BaseQuery[int64]
}
