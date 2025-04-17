package psychoclient

import (
	"fmt"

	"github.com/gogf/gf/container/gpool"
	"github.com/sicko7947/sickocommon"
)

// SessionPool interface allows operating a psycho client session pool
type SessionPool2 interface {
	Pop() Session
	Push(Session)
	Size() int
	Close()
}

type sessionPool2 struct {
	pool *gpool.Pool
}

func NewPool2(limit int, proxyGroup []string) SessionPool {
	pool := gpool.New(20000, func() (interface{}, error) {
		sesh, err := NewSession(&SessionBuilder{
			Proxy: sickocommon.GetProxy(proxyGroup).String(),
		})
		if err != nil {
			return nil, err.Error
		}

		return sesh, nil
		// return gtcp.NewConn("www.baidu.com:80")
	}, func(i interface{}) {
		fmt.Println("expired")
	})

	// pool := gqueue.New(limit)
	return &sessionPool2{pool: pool}
}

func (p *sessionPool2) Size() int {
	return p.pool.Size()
}

func (p *sessionPool2) Close() {
	p.pool.Close()
}
func (p *sessionPool2) Pop() Session {
	if s, _ := p.pool.Get(); s != nil {
		p.pool.Put(s)
		return s.(Session)
	}
	return nil
}

func (p *sessionPool2) Push(s Session) {
	defer recover()
	p.pool.Put(p)
	// p.pool.Push(s)
}
