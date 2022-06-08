package web

import (
	"fmt"
	"time"
)

type UnixTime time.Time

func (t UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", time.Time(t).UnixMilli())), nil
}

type ResponseModel struct {
	Data      interface{}  `json:"data"`
	Error     *interface{} `json:"error,omitempty"`
	Timestamp UnixTime     `json:"timestamp"`
}

func (r ResponseModel) Build() ResponseModel {
	r.Timestamp = UnixTime(time.Now())
	return r
}

type PageModel struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
}

type PageResult[T any] struct {
	List []T       `json:"list"`
	Page PageModel `json:"page,omitempty"`
}
