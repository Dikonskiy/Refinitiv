package main

import (
	"Refinitiv/internal/app"
	"Refinitiv/internal/config"
	"Refinitiv/internal/handlers"
	"Refinitiv/internal/models"
	"Refinitiv/internal/repository"
	"fmt"

	"github.com/gorilla/mux"
)

var (
	Repo *repository.Repository
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

	Repo = repository.NewRepository()
	Hand = handlers.NewHandlers(Repo)
	App = app.NewApplication()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/TokenManagement/TokenManagement.svc/REST/Anonymous/TokenManagement_1/{request}", Hand.MainHandler)

	App.StartServer(r, Cnfg)

}
