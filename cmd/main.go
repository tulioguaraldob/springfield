package main

import (
	"github.com/TulioGuaraldoB/springfield/internal/adapters/api"
	"github.com/TulioGuaraldoB/springfield/internal/config/env"
	"github.com/TulioGuaraldoB/springfield/utils/log/infologger"
)

func main() {
	infologger.CreateLog()

	env.GetConfigEnv()

	server := api.New()
	server.Run()
}
