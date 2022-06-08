package controller

import (
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/czyhome/go-pulsar-manager/service"
	"github.com/czyhome/go-pulsar-manager/web"
	"github.com/gin-gonic/gin"
)

func ClusterSearch(c *gin.Context) {
	list := []map[string]interface{}{
		{
			"id":   1,
			"name": "cluster1",
			"url":  "url1",
		},
		{
			"id":   2,
			"name": "cluster2",
			"url":  "url1",
		},
	}
	pageResult := web.PageResult[map[string]interface{}]{Page: web.PageModel{PageIndex: 1, PageSize: 10, Total: 10}, List: list}
	web.Context{Context: c}.OK(web.ResponseModel{Data: pageResult}.Build())

}

func ClusterAdd(c *gin.Context) {
	input := model.ClusterPO{}
	err := c.Bind(&input)
	web.CheckError(err)
	service.Cluster{}.Create(input)
}

func ClusterEdit(c *gin.Context) {

	//web.Response{Context: c,
	//	Data: map[string]interface{}{"status": "success"},
	//}.Build()

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
