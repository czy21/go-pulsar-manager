package util

import (
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
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
	var errMsg string
	if res.IsError() {
		errMsg = string(res.Body())
	}
	logParam := gin.LogFormatterParams{StatusCode: res.StatusCode(), Method: http.MethodGet, Path: url, ClientIP: h.BaseURL}
	log.Printf("  |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		logParam.StatusCodeColor(), logParam.StatusCode, logParam.ResetColor(),
		res.Time(),
		logParam.ClientIP,
		logParam.MethodColor(), logParam.Method, logParam.ResetColor(),
		logParam.Path,
		errMsg)
	err = h.JSONUnmarshal(res.Body(), v)
	exception.Check(err)
}
