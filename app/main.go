package main

import (
	// "encoding/base32"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "github.com/pquerna/otp"
	// "github.com/pquerna/otp/totp"
	"github.com/sauravkuila/portfolio-worth/api"
)

func main() {
	log.Println("Portfolio Valuation Project")
	// totp := generatePassCode("DFVGOUJ4T2MW356CCP5ZR7RAGQ")
	// fmt.Println(totp)
	if err := api.StartServer(); err != nil {
		log.Fatal("failed to start portfolio-worth server", err.Error())
	}
	// api.ReachAngel()
	// api.ReachZerodha()

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
