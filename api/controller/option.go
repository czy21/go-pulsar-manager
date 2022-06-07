package controller

import (
	"github.com/czyhome/go-pulsar-manager/web"
	"github.com/gin-gonic/gin"
)

func OptionList(c *gin.Context) {

	web.Response{Context: c,
		Data: map[string]interface{}{"status": "success"},
	}.Build()

}

func OptionController(r *gin.Engine) {

	v1 := r.Group("/option")
	{
		v1.POST("/query", OptionList)
	}

}
