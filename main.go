package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joshuabl97/product-api/handlers"
)

func main() {

	// instantiate logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProducts(l)

	// registering the handlers on the serve mux (sm)
	sm := mux.NewRouter()

	get := sm.Methods("GET").Subrouter()
	get.HandleFunc("/products", ph.GetProducts)

	put := sm.Methods("PUT").Subrouter()
	put.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProducts)

	post := sm.Methods("POST").Subrouter()
	post.HandleFunc("/products", ph.AddProduct)

	// create a new server
	s := http.Server{
		Addr:         ":8080",           // configure the bind address
		Handler:      sm,                // set the default handler
		IdleTimeout:  120 * time.Second, // max duration to wait for the next request when keep-alives are enabled
		ReadTimeout:  5 * time.Second,   // max duration for reading the request
		WriteTimeout: 10 * time.Second,  // max duration before returning the request
		ErrorLog:     l,                 // set the logger for the server
	}

	// this go function starts the server
	// when the function is done running, that means we need to shutdown the server
	// we can do this by killing the program, but if there are requests being processed
	// we want to give them time to complete
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// sending kill and interrupt signals to os.Signal channel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// does not envoke 'graceful shutdown' unless the signalChannel is closed
	sig := <-sigChan

	l.Println("Received terminate, graceful shutdown", sig)

	// this timeoutContext allows the server 30 seconds to complete all requests (if any) before shutting down
	timeoutCtx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := s.Shutdown(timeoutCtx)
	if err != nil {
		log.Fatal("We wanted to shut down anyway")
	}
}
