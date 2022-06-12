package controller

import (
	"fmt"
	"github.com/czyhome/go-pulsar-manager/constant"
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/czyhome/go-pulsar-manager/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func adminProxy(c *gin.Context) {
	input := model.EnvironmentQuery{}
	err := c.Bind(&input)
	exception.Check(err)
	remote, err := url.Parse(input.Url)
	exception.Check(err)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = fmt.Sprintf("/admin/v2/%s", c.Param("proxyPath"))
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func AdminController(r *gin.Engine) {
	r.Any(constant.PulsarRestPath+"/*proxyPath", adminProxy)
}
