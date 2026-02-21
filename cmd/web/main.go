package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grzegab/calendar/internal/config"
	"github.com/grzegab/calendar/internal/handlers"
	"github.com/grzegab/calendar/internal/ws"
)

// init is responsible for setting up CLI variables
func init() {
	// load variables from command line
	// e.g. go run main.go -config=app.yaml -debug
	flag.StringVar(&config.AppConfig.ConfigFile, "config", "config.yaml", "config file path")
	flag.BoolVar(&config.AppConfig.Debug, "debug", false, "should debug be enabled")
}

func main() {
	t := time.Now()
	fmt.Printf("[%s] calendar is starting out ...\n", t.Format("2006-01-02 15:04:05"))

	flag.Parse()

	// load config for k8s
	if err := config.LoadConfig(); err != nil {
		fmt.Printf("failed to load config: %v\n", err)
	}

	// load profiler if needed
	if config.AppConfig.Debug {
		go func() {
			fmt.Printf("starting pprof on %s\n", config.AppConfig.PprofAddr)
			if err := http.ListenAndServe(config.AppConfig.PprofAddr, nil); err != nil {
				fmt.Printf("pprof server error: %v\n", err)
			}
		}()
	} else {
		fmt.Println("pprof disabled in non-dev environment")
	}

	h := ws.NewHub()
	go h.Run()

	// router for requests
	r := handlers.New(config.AppConfig.Origins, h)
	server := &http.Server{
		Addr:         config.AppConfig.Addr,
		Handler:      r.Handler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		fmt.Printf("starting wall on %s\n", config.AppConfig.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	log.Println("Shutdown signal received, shutting down...")

	h.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}
