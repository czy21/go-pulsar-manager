package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"os"
)

func Boot() {
	r := gin.Default()
	logFile := viper.GetString("log.file")
	if logFile != "" {
		f, _ := os.Create(logFile)
		gin.DefaultWriter = io.MultiWriter(f)
	}
	AdminController(r)
	ClusterController(r)
	_ = r.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))
}
