package main

import (
	"context"
	"github.com/cockscomb/tinyurl/web"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	server := web.NewServer()

	go func() {
		if err := server.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalln("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}
