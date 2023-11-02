package main

import (
	"Refinitiv/internal/app"
	"Refinitiv/internal/config"
	"Refinitiv/internal/handlers"
	"Refinitiv/internal/models"
	"fmt"

	"github.com/gorilla/mux"
)

var (
	Cnfg *models.Config
	Hand *handlers.Handlers
	App  *app.Application
)

func init() {
	var err error
	Cnfg, err = config.InitConfig("config.json")
	if err != nil {
		fmt.Println("Failed to initialize the Configuration", err)
		return
	}

	Hand = handlers.NewHandlers()
	App = app.NewApplication()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/TokenManagement/TokenManagement.svc/REST/Anonymous/TokenManagement_1/CreateServiceToken_1", Hand.CreateServiceTokenHandler)

	App.StartServer(r, Cnfg)
}
