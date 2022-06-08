package web

import (
	"fmt"
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"

	"time"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case exception.Model:
					c.JSON(http.StatusOK, ResponseModel{Error: &err})
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

func DateValidator(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	fmt.Println(date, ok)
	return true
}
