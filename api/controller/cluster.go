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

func ClusterAdd(c *gin.Context) {

	web.Response{Context: c,
		Data: map[string]interface{}{"status": "success"},
	}.Build()

}

func ClusterEdit(c *gin.Context) {

	web.Response{Context: c,
		Data: map[string]interface{}{"status": "success"},
	}.Build()

}

func ClusterDel(c *gin.Context) {

	web.Response{Context: c,
		Data: map[string]interface{}{"status": "success"},
	}.Build()

}

func ClusterController(r *gin.Engine) {

	v1 := r.Group("/cluster")
	{
		v1.POST("/list", ClusterList)
		v1.POST("/add", ClusterAdd)
		v1.POST("/edit", ClusterEdit)
		v1.POST("/del", ClusterDel)
	}

}
