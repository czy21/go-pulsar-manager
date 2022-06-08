package model

type ClusterPO struct {
	BaseEntity[int64]
	Name        string `gorm:"column:name" json:"name"`
	Url         string `gorm:"column:url" json:"url"`
	Description string `gorm:"column:description" json:"description"`
}
