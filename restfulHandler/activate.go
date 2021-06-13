package restfulhandler

import (
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

	r.Response.WriteJsonExit(map[string]int64{
		"status": int64(code),
	})
}
