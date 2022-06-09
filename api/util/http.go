package util

import (
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/go-resty/resty/v2"
)

type HttpUtil struct {
	*resty.Client
}

func (HttpUtil) NewClient() HttpUtil {
	c := resty.New()
	return HttpUtil{Client: c}
}

func (h HttpUtil) Get(url string, v interface{}) {
	res, err := h.R().Get(url)
	exception.Check(err)
	err = h.JSONUnmarshal(res.Body(), v)
	exception.Check(err)
}
