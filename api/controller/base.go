package controller

import (
	"fmt"
	"github.com/czyhome/go-pulsar-manager/web"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
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
	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		ResourceController(r1)
	}
	address := fmt.Sprintf(":%s", viper.GetString("server.port"))
	log.Printf("Listening and serving HTTP on %s\n", address)
	_ = http.ListenAndServe(address, r1)
}
