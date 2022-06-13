package controller

import (
	"fmt"
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func resourceProxy(c *gin.Context) {
	remote, err := url.Parse(fmt.Sprintf("http://127.0.0.1:%s", viper.GetString("server.port")))
	exception.Check(err)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func ResourceController(r *gin.Engine) {
	indexFile := fmt.Sprintf("%s/index.html", "dist")
	staticFile := fmt.Sprintf("%s/static/", "dist")
	r.NoRoute(func(c *gin.Context) {
		c.File(indexFile)
	})
	r.StaticFile("/", indexFile)
	r.Static("/static/", staticFile)
	r.Any("/api/*proxyPath", resourceProxy)
}
