package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"PelandoChallenge/pkg/producthandlers"
	"PelandoChallenge/pkg/service"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	log := logrus.New()
	addr := fmt.Sprintf(":%d", 8080)

	srvc, err := service.New(ctx)
	if err != nil {
		log.Errorf("Failed to start service: %v", err)
		return
	}

	handler := chi.NewRouter()
	{
		productHandler := producthandlers.NewProductHandler(srvc)

		handler.Route("/api/", func(r chi.Router) {
			r.Get("/product", productHandler.HandleProduct)
		})
	}

	server := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Errorf("Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Infof("Listening on %s", addr)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Errorf("ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
