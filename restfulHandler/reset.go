package restfulhandler

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/sicko7947/sicko-aio-auth/postgresql"
)

func reset(r *ghttp.Request) {
	key := r.Get("key").(string)
	email := r.Get("email").(string)
	discordID := r.Get("discordID").(string)

	code, err := postgresql.Reset(key, email, discordID)
	if err != nil {
		r.Response.WriteStatusExit(500)
		return
	}

	r.Response.WriteJsonExit(map[string]interface{}{
		"code":   int64(code),
		"result": postgresql.STATUSMAP[code],
	})
}
