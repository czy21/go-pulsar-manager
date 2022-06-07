package bootstrap

import "github.com/gin-gonic/gin"
import controller "github.com/czyhome/go-pulsar-manager/controller"

func bootRouter() {
	r := gin.Default()
	controller.AdminController(r)
	controller.ClusterController(r)
	_ = r.Run(":8080")
}
