package web

import (
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case exception.Model:
					c.JSON(http.StatusOK, model.ResponseModel{Error: &err})
					break
				}
				panic(err)
			}
		}()
		c.Next()
	}
}

func ResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
