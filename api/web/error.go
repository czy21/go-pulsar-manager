package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case ExceptionModel:
					c.JSON(http.StatusOK, ResponseModel{Error: &err})
					break
				}
				panic(err)
			}
		}()
		c.Next()
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type ExceptionModel struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewException(message string) ExceptionModel {
	panic(ExceptionModel{Message: message})
}
