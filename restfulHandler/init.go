package restfulhandler

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func portInUse(port int) bool {
	checkStatement := fmt.Sprintf(`netstat -anp | grep -q %d ; echo $?`, port)
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	// log.Println(output, string(output)) ==> [48 10] 0 或 [49 10] 1
	result, err := strconv.Atoi(strings.TrimSuffix(string(output), "\n"))
	if err != nil {
		return true
	}
	if result == 0 {
		return true
	}
	return false
}

func init() {
	fmt.Println(portInUse(28888))
	if !portInUse(28888) {
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
