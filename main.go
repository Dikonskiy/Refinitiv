package main

import (
	"Refinitiv/internal/app"
	"Refinitiv/internal/config"
	"fmt"
)

func main() {
	cnfg, err := config.InitConfig("config.json")
	if err != nil {
		fmt.Println("Failed to initialize the Configuration", err)
		return
	}

	app := app.NewApplication()

	app.StartServer(cnfg.ListenPort, cnfg.Route)

}
