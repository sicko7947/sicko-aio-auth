package postgresql

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/names"
)

var eg *xorm.EngineGroup

const (
	host     = "131.153.76.114"
	port     = 5432
	user     = "sickomcd"
	password = "tWn8sSheAWHC"
	dbname   = "keymainv2"
)

func init() {
	var err error

	conns := []string{
		"postgres://sickomcd:tWn8sSheAWHC@131.153.76.114:5432/keymainv2?sslmode=disable;", // 第一个默认是master
		"postgres://sickomcd:tWn8sSheAWHC@131.153.76.114:5432/keymainv3?sslmode=disable;", // 第二个开始都是slave
	}
	eg, err = xorm.NewEngineGroup("postgres", conns, xorm.LeastConnPolicy())
	if err != nil {
		log.Fatal(err)
	}

	err = eg.Ping()

	eg.SetMaxIdleConns(50)
	eg.SetMaxOpenConns(5)
	eg.SetMapper(names.SameMapper{})
	eg.Sync2(new(successTable), new(productDetail), new(keyMain), new(keyDetails))

	if err != nil {
		log.Fatal(err)
	}
}
