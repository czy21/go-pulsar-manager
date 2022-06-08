package model

type Tabler interface {
	TableName() string
}

type PageModel struct {
	PageIndex int   `json:"pageIndex"`
	PageSize  int   `json:"pageSize"`
	Total     int64 `json:"total"`
}

type PageResult[T any] struct {
	List []T       `json:"list"`
	Page PageModel `json:"page,omitempty"`
}

type BaseEntity[TID any] struct {
	Id         TID       `gorm:"column:id" json:"id"`
	CreateTime *UnixTime `gorm:"column:create_time;default:null" json:"createTime"`
	CreateUser string    `gorm:"column:create_user" json:"createUser"`
	UpdateTime *UnixTime `gorm:"column:update_time;default:null" json:"updateTime"`
	UpdateUser string    `gorm:"column:update_user" json:"updateUser"`
}

type BaseQuery[TID any] struct {
	Id   TID    `json:"id"`
	Name string `json:"name"`
	PageModel
}

type SimpleItemModel[T any] struct {
	Label    string                 `json:"label"`
	Value    T                      `json:"value"`
	Extra    map[string]interface{} `json:"extra,omitempty"`
	Children []SimpleItemModel[T]   `json:"children,omitempty"`
}