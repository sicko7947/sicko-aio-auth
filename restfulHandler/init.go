package restfulhandler

import (
	"fmt"
	"os/exec"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func PortInUse(port int) bool {
	checkStatement := fmt.Sprintf("lsof -i:%d ", port)
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	return len(output) > 0
}

func init() {
	if !PortInUse(28888) {
		s := g.Server()

		s.Group("/", func(group *ghttp.RouterGroup) {
			group.POST("/auth/activate", activate)
			group.POST("/auth/reset", reset)
		})
		s.EnableHTTPS(`/root/auth/auth.sickoaio.com.crt`, `/root/auth/auth.sickoaio.com.key`)
		s.SetHTTPSPort(28888)
		s.Run()
	}
}

// func MiddlewareCORS(r *ghttp.Request) {
// 	r.Response.CORSDefault()
// 	r.Middleware.Next()
// }
