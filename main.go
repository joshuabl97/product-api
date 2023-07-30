package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	redoc "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/joshuabl97/product-api/handlers"
)

func main() {

	// instantiate logger
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	// setting timezone
	loc, err := time.LoadLocation("Etc/Greenwich")
	if err != nil {
		l.Error().Msg("Couldn't determine timezone, using local machine time")
	} else if err == nil {
		time.Local = loc
	}

	// make the logs look pretty
	l = l.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	// create a custom logger that wraps the zerolog.Logger we instantiated/customized above
	errorLog := &zerologLogger{l}

	// create the handlers
	ph := handlers.NewProducts(&l)

	// registering the handlers on the serve mux (sm)
	sm := mux.NewRouter()

	get := sm.Methods("GET").Subrouter()
	get.HandleFunc("/products", ph.GetProducts)
	get.HandleFunc("/products/{id:[0-9]+}", ph.GetProduct)

	del := sm.Methods("DELETE").Subrouter()
	del.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	put := sm.Methods("PUT").Subrouter()
	put.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProduct)
	// the handler func (ph.UpdateProducts) will only be envoked if the middleware passes
	put.Use(ph.MiddlewareProductValidation)

	post := sm.Methods("POST").Subrouter()
	post.HandleFunc("/products", ph.AddProduct)
	post.Use(ph.MiddlewareProductValidation)

	opts := redoc.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := redoc.Redoc(opts, nil)

	// displays swagger page
	get.Handle("/docs", handlers.LoggingMiddleware(sh, &l))
	// the middleware.Redoc handler requires access to the swagger file
	// as defined above in the middleware.RedocOpts
	get.Handle("/swagger.yaml", handlers.LoggingMiddleware(http.FileServer(http.Dir("./")), &l))

	// create a new server
	s := http.Server{
		Addr:         ":8080",                  // configure the bind address
		Handler:      sm,                       // set the default handler
		IdleTimeout:  120 * time.Second,        // max duration to wait for the next request when keep-alives are enabled
		ReadTimeout:  5 * time.Second,          // max duration for reading the request
		WriteTimeout: 10 * time.Second,         // max duration before returning the request
		ErrorLog:     log.New(errorLog, "", 0), // set the logger for the server
	}

	// this go function starts the server
	// when the function is done running, that means we need to shutdown the server
	// we can do this by killing the program, but if there are requests being processed
	// we want to give them time to complete
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal().Err(err)
		}
	}()

	// sending kill and interrupt signals to os.Signal channel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// does not invoke 'graceful shutdown' unless the signalChannel is closed
	<-sigChan

	l.Info().Msg("Received terminate, graceful shutdown")

	// this timeoutContext allows the server 30 seconds to complete all requests (if any) before shutting down
	timeoutCtx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = s.Shutdown(timeoutCtx)
	if err != nil {
		l.Fatal().Msg("We wanted to shut down anyway")
	}
}

// custom logger type that wraps zerolog.Logger
type zerologLogger struct {
	logger zerolog.Logger
}

// implement the io.Writer interface for our custom logger.
func (l *zerologLogger) Write(p []byte) (n int, err error) {
	l.logger.Error().Msg(string(p))
	return len(p), nil
}
