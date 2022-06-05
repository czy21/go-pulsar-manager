package main

import (
	"github.com/czyhome/go-pulsar-manager/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.Admin(r)
	_ = r.Run(":8080")
}
