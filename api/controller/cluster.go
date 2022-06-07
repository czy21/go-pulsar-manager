package controller

import (
	"github.com/czyhome/go-pulsar-manager/web"
	"github.com/gin-gonic/gin"
)

func ClusterList(c *gin.Context) {

	web.Response{Context: c,
		Data: map[string]interface{}{"status": "success"},
	}.Build()

}

func ClusterController(r *gin.Engine) {

	v1 := r.Group("/cluster")
	{
		v1.POST("/list", ClusterList)
	}

}
