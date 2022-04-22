package main

import (
	"github.com/TulioGuaraldoB/springfield/config/env"
	"github.com/TulioGuaraldoB/springfield/server"
	"github.com/TulioGuaraldoB/springfield/utils/log/infologger"
)

func main() {
	infologger.CreateLog()

	env.GetConfigEnv()

	server := server.NewServer()
	server.RunServer()
}
