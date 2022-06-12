package controller

import (
	"github.com/gin-gonic/gin"
)

func resourceProxy(c *gin.Context) {
	//remote, err := url.Parse("http://127.0.0.1:8080")
	//exception.Check(err)
	//proxy := httputil.NewSingleHostReverseProxy(remote)
	//proxy.Director = func(req *http.Request) {
	//	req.Header = c.Request.Header
	//	req.Host = remote.Host
	//	req.URL.Scheme = remote.Scheme
	//	req.URL.Host = remote.Host
	//	req.URL.Path = c.Param("proxyPath")
	//}
	//proxy.ServeHTTP(c.Writer, c.Request)
}

func ResourceController(r *gin.Engine) {
	r.Any("/api/*proxyPath", resourceProxy)
}
