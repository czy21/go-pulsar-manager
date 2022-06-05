package controller

import (
	"github.com/gin-gonic/gin"
)

func adminProxy(c *gin.Context) {
	//remote, err := url.Parse("http://baidu.com")
	//if err != nil {
	//	panic(err)
	//}
	//proxy := httputil.NewSingleHostReverseProxy(remote)
	//proxy.Director = func(req *http.Request) {
	//	req.Header = c.Request.Header
	//	req.Host = remote.Host
	//	req.URL.Scheme = remote.Scheme
	//	req.URL.Host = remote.Host
	//req.URL.Path = c.Param("proxyPath")
	//}
	//
	//proxy.ServeHTTP(c.Writer, c.Request)
}

func Admin(r *gin.Engine) {
	r.Any("/admin/v2/*proxyPath", adminProxy)
}
