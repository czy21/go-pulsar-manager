package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func Boot() {
	gin.DefaultWriter = log.Writer()
	r := gin.Default()
	EnvironmentController(r)
	AdminController(r)
	ClusterController(r)
	TenantController(r)
	_ = r.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))
}
