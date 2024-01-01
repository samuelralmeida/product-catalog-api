package api

// https://medium.com/inside-picpay/abstraindo-bibliotecas-web-em-aplica%C3%A7%C3%B5es-go-764ebd2ba200

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/samuelralmeida/product-catalog-api/env"
)

const TIMEOUT = 30 * time.Second

// Start a new http server with graceful shutdown and default parameters
func Start(config *env.Config, handler http.Handler) error {

	// n := negroni.New()
	// n.Use(negroni.HandlerFunc(MyMiddleware))
	// n.Use(negroni.Wrap(middlewares.SetUser()))
	// n.UseHandler(handler)

	srv := &http.Server{
		ReadTimeout:  TIMEOUT,
		WriteTimeout: TIMEOUT,
		Addr:         config.Server.Port,
		// Handler:      n,
		Handler: handler,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
	)
	defer stop()
	errShutdown := make(chan error, 1)
	go shutdown(srv, ctx, errShutdown)

	log.Printf("Current service listening on port %s\n", config.Server.Port)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	err = <-errShutdown
	if err != nil {
		return err
	}
	return nil

}

func shutdown(server *http.Server, ctxShutdown context.Context, errShutdown chan error) {
	<-ctxShutdown.Done()

	ctxTimeout, stop := context.WithTimeout(context.Background(), TIMEOUT)
	defer stop()

	err := server.Shutdown(ctxTimeout)
	switch err {
	case nil:
		errShutdown <- nil
	case context.DeadlineExceeded:
		errShutdown <- errors.New("forcing closing the server")
	default:
		errShutdown <- errors.New("forcing closing the server")
	}
}
