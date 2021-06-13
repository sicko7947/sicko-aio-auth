package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sicko7947/sicko-aio-auth/grpcHandler"
	restfulhandler "github.com/sicko7947/sicko-aio-auth/restfulHandler"
)

func main() {
	port, _ := strconv.ParseInt(os.Args[1], 10, 64)
	portString := fmt.Sprintf(":%v", port)
	go grpcHandler.StargrpcServer(portString)
	restfulhandler.BackendServer()
}
