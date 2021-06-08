package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sicko7947/sicko-aio-auth/grpcHandler"
)

func main() {
	port, _ := strconv.ParseInt(os.Args[1], 10, 64)
	portString := fmt.Sprintf(":%v", port)
	grpcHandler.StargrpcServer(portString)
}
