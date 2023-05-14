package main

import (
	"github.com/writeas/web-core/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Info("Shutting down...")
		SaveCounts()
		log.Info("Done.")
		os.Exit(0)
	}()

	err := Serve()
	if err != nil {
		log.Error("Unable to serve: %s", err)
		os.Exit(1)
	}
}
