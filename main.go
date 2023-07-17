package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joshuabl97/product-api/handlers"
)

func main() {

	// instantiate logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	// **this is just for example, but you can inject the handler directly into the ServeMux.Handle()
	hh := handlers.NewHello(l)

	// registering the handlers on the serve mux (sm)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	// **example of registering a route to the serve mux by passing the http.Handler in directly
	sm.Handle("/goodbye", handlers.NewGoodbye(l))

	// create a new server
	s := http.Server{
		Addr:         ":8080",           // configure the bind address
		Handler:      sm,                // set the default handler
		IdleTimeout:  120 * time.Second, // max duration to wait for the next request when keep-alives are enabled
		ReadTimeout:  1 * time.Second,   // max duration for reading the request
		WriteTimeout: 1 * time.Second,   // max duration before returning the request
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
