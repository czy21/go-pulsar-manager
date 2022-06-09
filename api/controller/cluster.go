package controller

import (
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/service"
	"github.com/czyhome/go-pulsar-manager/web"
	"github.com/gin-gonic/gin"
)

func ClusterSearch(c *gin.Context) {
	input := model.ClusterQuery{}
	err := c.Bind(&input)
	exception.Check(err)
	list := service.Cluster{}.FindList(input)
	web.Context{Context: c}.OK(model.ResponseModel{Data: list}.Build())
}

func ClusterController(r *gin.Engine) {

	v1 := r.Group("/cluster")
	{
		v1.POST("/search", ClusterSearch)
	}

}
