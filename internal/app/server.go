package app

import (
	"Refinitiv/internal/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) StartServer(router http.Handler, config *models.Config) {
	server := &http.Server{
		Addr:         ":" + config.ListenPort,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
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
