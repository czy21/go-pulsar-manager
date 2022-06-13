package controller

import (
	"github.com/czyhome/go-pulsar-manager/web"
	"github.com/gin-gonic/gin"
)

func ApiEngine() *gin.Engine {
	r := gin.New()
	r.Use(web.ErrorHandle())
	OptionController(r)
	EnvironmentController(r)
	AdminController(r)
	ClusterController(r)
	TenantController(r)
	return r
}
