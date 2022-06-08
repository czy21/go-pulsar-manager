package controller

import (
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/service"
	"github.com/czyhome/go-pulsar-manager/web"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ClusterSearch(c *gin.Context) {
	input := model.ClusterQuery{}
	err := c.Bind(&input)
	exception.Check(err)
	pageResult := service.Cluster{}.FindList(input)
	web.Context{Context: c}.OK(model.ResponseModel{Data: pageResult}.Build())
}

func ClusterAdd(c *gin.Context) {
	input := model.ClusterPO{}
	err := c.Bind(&input)
	exception.Check(err)
	service.Cluster{}.Create(input)
}

func ClusterEdit(c *gin.Context) {

	//web.Response{Context: c,
	//	Data: map[string]interface{}{"status": "success"},
	//}.Build()
	web.Context{Context: c}.OK(model.ResponseModel{Data: viper.Get("name")}.Build())
}

func ClusterDel(c *gin.Context) {

	//web.Response{Context: c,
	//	Data: map[string]interface{}{"status": "success"},
	//}.Build()

}

func ClusterController(r *gin.Engine) {

	v1 := r.Group("/cluster")
	{
		v1.POST("/search", ClusterSearch)
		v1.POST("/add", ClusterAdd)
		v1.POST("/edit", ClusterEdit)
		v1.POST("/del", ClusterDel)
	}

}
