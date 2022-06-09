package controller

import (
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/service"
	"github.com/czyhome/go-pulsar-manager/web"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func EnvironmentSearch(c *gin.Context) {
	input := model.EnvironmentQuery{}
	err := c.Bind(&input)
	exception.Check(err)
	pageResult := service.Environment{}.FindList(input)
	web.Context{Context: c}.OK(model.ResponseModel{Data: pageResult}.Build())
}

func EnvironmentAdd(c *gin.Context) {
	input := model.EnvironmentPO{}
	err := c.Bind(&input)
	exception.Check(err)
	service.Environment{}.Create(input)
}

func EnvironmentEdit(c *gin.Context) {

	//web.Response{Context: c,
	//	Data: map[string]interface{}{"status": "success"},
	//}.Build()
	web.Context{Context: c}.OK(model.ResponseModel{Data: viper.Get("name")}.Build())
}

func EnvironmentDel(c *gin.Context) {

	//web.Response{Context: c,
	//	Data: map[string]interface{}{"status": "success"},
	//}.Build()

}

func EnvironmentController(r *gin.Engine) {

	v1 := r.Group("/environment")
	{
		v1.POST("/search", EnvironmentSearch)
		v1.POST("/add", EnvironmentAdd)
		v1.POST("/edit", EnvironmentEdit)
		v1.POST("/del", EnvironmentDel)
	}

}
