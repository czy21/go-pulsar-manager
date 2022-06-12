package controller

import (
	"fmt"
	"github.com/czyhome/go-pulsar-manager/web"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func Boot() {
	gin.DefaultWriter = log.Writer()
	r := gin.Default()
	r.Use(web.ErrorHandle())
	OptionController(r)
	EnvironmentController(r)
	AdminController(r)
	ClusterController(r)
	TenantController(r)
	_ = r.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))
}
