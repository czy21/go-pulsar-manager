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
	r1 := gin.Default()
	r1.Use(web.ErrorHandle())
	OptionController(r1)
	EnvironmentController(r1)
	AdminController(r1)
	ClusterController(r1)
	TenantController(r1)
	ResourceController(r1)
	_ = r1.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))
}
