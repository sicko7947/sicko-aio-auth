package restfulhandler

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)

		group.POST("/auth/reset", reset)
	})
	s.EnableHTTPS(`/root/auth/auth.sickoaio.com.crt`, `/root/auth/auth.sickoaio.com.key`)
	s.SetPort(28888)
	s.Run()
}

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
