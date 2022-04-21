package main

import (
	"github.com/TulioGuaraldoB/springfield/server"
)

func main() {
	server := server.NewServer()
	server.RunServer()
}
