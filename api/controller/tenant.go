package controller

import "github.com/gin-gonic/gin"

func TenantSearch(c *gin.Context) {
	//web.Response{Context: c,
	//	Data: map[string]interface{}{"status": "success"},
	//}.Build()

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
