package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Boot() {
	engine := gin.Default()
	AdminController(engine)
	ClusterController(engine)
	_ = engine.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))
}
