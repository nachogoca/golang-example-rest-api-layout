package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nachogoca/golang-example-rest-api-layout/internal/middlewares"
	"github.com/nachogoca/golang-example-rest-api-layout/internal/stores"
	"github.com/nachogoca/golang-example-rest-api-layout/internal/transports"
	"github.com/nachogoca/golang-example-rest-api-layout/internal/usecases"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	// Init service, usecase and transport layers
	// Clean code architecture is used here
	store, err := stores.NewArticles()
	if err != nil {
		logrus.WithError(err).Fatal("could not init store")
	}
	defer store.Close()

	usecase := usecases.NewArticles(store)
	transport := transports.NewArticles(usecase)

	// Init router
	r := mux.NewRouter()
	s := r.PathPrefix("/articles").Subrouter()
	s.HandleFunc("", transport.GetAll).Methods("GET")
	s.HandleFunc("", transport.Create).Methods("POST")
	s.HandleFunc("/{id}", transport.GetOne).Methods("GET")
	s.HandleFunc("/{id}", transport.Update).Methods("PUT")
	s.HandleFunc("/{id}", transport.Delete).Methods("DELETE")

	// Set middlewares
	r.Use(middlewares.RequestID)
	r.Use(middlewares.Logging)

	// Init server with timeouts
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// Start server in goroutine
	go func() {
		logrus.Warn("Starting server")
		if err := srv.ListenAndServe(); err != nil {
			logrus.WithError(err).Error("Got server error")
		}
	}()

	// Handle graceful shutdowns via SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Wait for requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Graceful shutdown
	srv.Shutdown(ctx)
	logrus.Warn("Shutting down gracefully")
	os.Exit(0)
}
