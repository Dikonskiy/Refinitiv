package main

import (
	"Refinitiv/internal/app"
	"Refinitiv/internal/config"
	"Refinitiv/internal/handlers"
	"Refinitiv/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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
	request := models.CreateServiceTokenRequest{
		CreateServiceTokenRequest1: struct {
			ApplicationID string `json:"ApplicationID"`
			Username      string `json:"Username"`
			Password      string `json:"Password"`
		}{
			ApplicationID: "1",
			Username:      "Dias",
			Password:      "dias111",
		},
	}

	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(requestJSON))

	r := mux.NewRouter()

	r.HandleFunc("/api/TokenManagement/TokenManagement.svc/REST/Anonymous/TokenManagement_1/CreateServiceToken_1", Hand.CreateServiceTokenHandler)

	App.StartServer(r, Cnfg)

	url := "http://localhost:8080/api/TokenManagement/TokenManagement.svc/REST/Anonymous/TokenManagement_1/CreateServiceToken_1"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
		return
	}
}
