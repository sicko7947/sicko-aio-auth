package restfulhandler

import (
	"fmt"

	"github.com/gogf/gf/net/ghttp"
	"github.com/sicko7947/sicko-aio-auth/postgresql"
)

func activate(r *ghttp.Request) {
	key := r.Get("key").(string)
	email := r.Get("email").(string)
	discordID := r.Get("discordID").(string)

	code, err := postgresql.Activate(key, email, discordID)
	if err != nil {
		r.Response.WriteStatusExit(500)
		return
	}

	r.Response.WriteJsonExit(map[string]string{
		"code":   fmt.Sprint(code),
		"result": postgresql.STATUSMAP[code],
	})
}
