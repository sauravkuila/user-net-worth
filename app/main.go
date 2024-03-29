package main

import (
	// "encoding/base32"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sauravkuila/portfolio-worth/api"
	"github.com/sauravkuila/portfolio-worth/pkg/config"
)

func main() {
	log.Println("Portfolio Valuation Project")
	//initialize config
	filename := "local"
	if _, ok := os.LookupEnv("HOST"); ok {
		filename = "server"
	}
	config.Load(filename)

	if err := api.StartServer(); err != nil {
		log.Fatal("failed to start portfolio-worth server", err.Error())
	}

	// Wait for interrupt signal to gracefully shutdown the server with a timeouts
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down Server ...")
	wrapUp()

	log.Println("Server shutdown gracefully")
}

func wrapUp() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := api.ShutDownServer(ctx); err != nil {
		log.Fatal("Server Shutdown with error:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("server shutdown timeout. force exit.")
		return
	}
}
