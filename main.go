package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/baldore/building-microservices-with-go/handlers"
	"github.com/gorilla/mux"
)

func setupServer(l *log.Logger) *mux.Router {
	hh := handlers.NewHello(l)
	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()
	sm.NotFoundHandler = handlers.Handle404

	sm.HandleFunc("/hello", hh.SayHello).Methods("GET")

	sm.HandleFunc("/products", ph.GetProducts).Methods("GET")

	return sm
}

func main() {
	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)

	sm := setupServer(l)

	s := &http.Server{
		Addr:    ":9090",
		Handler: sm,

		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// TODO: What are go func (goroutines)?
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// TODO: What are channels?
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
