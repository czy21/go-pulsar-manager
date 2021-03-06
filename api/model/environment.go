package model

type EnvironmentPO struct {
	BaseEntity[int64]
	TrackEntity
	Name        string `gorm:"column:name" json:"name"`
	Url         string `gorm:"column:url" json:"url"`
	Description string `gorm:"column:description" json:"description"`
}

func (EnvironmentPO) TableName() string {
	return "environment"
}

type EnvironmentDTO struct {
	BaseEntity[int64]
	TrackEntity
	Name        string `gorm:"column:name" json:"name"`
	Url         string `gorm:"column:url" json:"url"`
	Description string `gorm:"column:description" json:"description"`
}

type EnvironmentQuery struct {
	BaseQuery[int64]
	Url string `json:"url"`
}
