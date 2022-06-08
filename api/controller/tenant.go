package controller

import (
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/service"
	"github.com/czyhome/go-pulsar-manager/web"
	"github.com/gin-gonic/gin"
)

func TenantSearch(c *gin.Context) {
	input := model.TenantQuery{}
	err := c.Bind(&input)
	exception.Check(err)
	web.Context{Context: c}.OK(model.ResponseModel{Data: service.Tenant{}.FindList(input)}.Build())
}

func TenantAdd(c *gin.Context) {

	//web.Response{Context: c,
	//	Data: map[string]interface{}{"status": "success"},
	//}.Build()

}

func TenantEdit(c *gin.Context) {

	//web.Response{Context: c,
	//	Data: map[string]interface{}{"status": "success"},
	//}.Build()

}

func TenantDel(c *gin.Context) {
	//web.Response{Context: c,
	//	Data: map[string]interface{}{"status": "success"},
	//}.Build()

}

func TenantController(r *gin.Engine) {

	v1 := r.Group("/tenant")
	{
		v1.POST("/search", TenantSearch)
		v1.POST("/add", TenantAdd)
		v1.POST("/edit", TenantEdit)
		v1.POST("/del", TenantDel)
	}

}
