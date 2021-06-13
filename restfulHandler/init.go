package restfulhandler

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func BackendServer(port int64) {
	if port != 26600 {
		return
	}

	s := g.Server()

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.POST("/auth/activate", activate)
		group.POST("/auth/reset", reset)
	})
	s.EnableHTTPS(`/root/auth/auth.sickoaio.com.crt`, `/root/auth/auth.sickoaio.com.key`)
	s.SetHTTPSPort(28888)
	s.Run()
}

// func MiddlewareCORS(r *ghttp.Request) {
// 	r.Response.CORSDefault()
// 	r.Middleware.Next()
// }
