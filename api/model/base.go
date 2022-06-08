package model

import "github.com/czyhome/go-pulsar-manager/web"

type Tabler interface {
	TableName() string
}

type BaseEntity[TID any] struct {
	Id         TID           `gorm:"column:id" json:"id"`
	CreateTime *web.UnixTime `gorm:"column:create_time;default:null" json:"createTime"`
	CreateUser string        `gorm:"column:create_user" json:"createUser"`
	UpdateTime *web.UnixTime `gorm:"column:update_time;default:null" json:"updateTime"`
	UpdateUser string        `gorm:"column:update_user" json:"updateUser"`
}

type BaseQuery[TID any] struct {
	Id   TID    `json:"id"`
	Name string `json:"name"`
	web.PageModel
}

type SimpleItemModel[T any] struct {
	Label    string                 `json:"label"`
	Value    T                      `json:"value"`
	Extra    map[string]interface{} `json:"extra,omitempty"`
	Children []SimpleItemModel[T]   `json:"children,omitempty"`
}
