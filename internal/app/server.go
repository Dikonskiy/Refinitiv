package app

import (
	"Refinitiv/internal/handlers"
	"Refinitiv/internal/quotes"
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
	Quotes *quotes.Quotes
	Token  *tokenizer.Tokenizer
	Hand   *handlers.Handlers
)

func init() {
	Token = tokenizer.NewTokenizer()
	Quotes = quotes.NewQuotes()
	Hand = handlers.NewHandlers(Token, Quotes)
}

func (a *Application) StartServer(listenPort, route string) {
	r := mux.NewRouter()

	subrouter := r.PathPrefix(route).Subrouter()

	subrouter.HandleFunc("/CreateServiceToken_1", Hand.CreateServiceTokenHandler)
	subrouter.HandleFunc("/ValidateServiceToken_1", Hand.ValidateServiceTokenHandler)
	subrouter.HandleFunc("/CreateImpersonationToken_1", Hand.CreateImpersonationTokenHandler)
	subrouter.HandleFunc("/CreateImpersonationToken_2", Hand.GenerateServiceAndImpersonationToken)
	subrouter.HandleFunc("/RetrieveItem_Request_3", applyMiddleware(Hand.GetQuotes, setTokenMiddleware(Token))).Methods("POST")

	server := &http.Server{
		Addr:         ":" + listenPort,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
	}

	quit := make(chan os.Signal, 1)

	go shutdown(quit)

	fmt.Println("Listening on port", listenPort, "...")
	log.Fatal(server.ListenAndServe())
}

func shutdown(quit chan os.Signal) {
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	fmt.Println("caught interrupt signal", s.String())
	os.Exit(0)
}

func setTokenMiddleware(tokenizer *tokenizer.Tokenizer) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username := "Dias"
			applicationID := "1"

			fmt.Println("Middleware: username =", username, "applicationID =", applicationID)
			fmt.Println("ServiceTokens map:", tokenizer.ServiceTokens)

			token, exists := tokenizer.ServiceTokens[applicationID][username]
			if !exists {
				fmt.Println("Token not found")
				http.Error(w, "Token not found", http.StatusUnauthorized)
				return
			}

			r.Header.Set("Authorization", token)
			r.Header.Set("ApplicationID", applicationID)

			next.ServeHTTP(w, r)
		})
	}
}

func applyMiddleware(h http.HandlerFunc, middleware mux.MiddlewareFunc) http.HandlerFunc {
	return middleware(h).ServeHTTP
}
