package app

import (
	"Refinitiv/internal/handlers"
	"Refinitiv/internal/models"
	"Refinitiv/internal/tokenizer"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

var (
	Token *tokenizer.Tokenizer
	Hand  *handlers.Handlers
)

func init() {
	Token = tokenizer.NewTokenizer()
	Hand = handlers.NewHandlers(Token)
}

func (a *Application) StartServer(config *models.Config) {
	r := mux.NewRouter()

	r.HandleFunc("/api/TokenManagement/TokenManagement.svc/REST/Anonymous/TokenManagement_1/CreateServiceToken_1", Hand.CreateServiceTokenHandler)
	r.HandleFunc("/api/TokenManagement/TokenManagement.svc/REST/Anonymous/TokenManagement_1/ValidateServiceToken_1", Hand.ValidateServiceTokenHandler)
	r.HandleFunc("/api/TokenManagement/TokenManagement.svc/REST/Anonymous/TokenManagement_1/CreateImpersonationToken_1", Hand.CreateImpersonationTokenHandler)
	r.HandleFunc("/api/TokenManagement/TokenManagement.svc/REST/Anonymous/TokenManagement_1/CreateImpersonationToken_2", Hand.GenerateServiceAndImpersonationToken)
	r.HandleFunc("/api/TokenManagement/TokenManagement.svc/REST/Anonymous/TokenManagement_1/RetrieveItem_Request_3", Hand.GetQuotes)

	server := &http.Server{
		Addr:         ":" + config.ListenPort,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
	}

	quit := make(chan os.Signal, 1)

	go shutdown(quit)

	fmt.Println("Listening on port", config.ListenPort, "...")
	log.Fatal(server.ListenAndServe())
}

func shutdown(quit chan os.Signal) {
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	fmt.Println("caught interrupt signal", s.String())
	os.Exit(0)
}
