package main

import (
	"DBIH/controller/module"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/gorilla/mux"
)

// encrypt the listKey and store it in postgreSQL
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/modify/{uid}", module.ModifyHandler).Methods("POST")
	router.HandleFunc("/get/{uid}/{page}", module.Getpage).Methods("GET")
	srv := &http.Server{
		Addr:    ":8006",
		Handler: router,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()
	<-sig

	log.Println("shutting down...")
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
	log.Println("shutdown complete")
}
