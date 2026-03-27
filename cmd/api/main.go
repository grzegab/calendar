package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github/grzegab/calendar/internal/app"
	"github/grzegab/calendar/internal/app/debug"
	"github/grzegab/calendar/internal/shared/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github/grzegab/calendar/internal/shared/infrastructure/db"
	"github/grzegab/calendar/internal/shared/ws"
)

// init is responsible for setting up CLI variables
func init() {
	// load variables from command line, e.g. go run main.go -config=.env
	flag.StringVar(&app.AppConfig.ConfigFile, "env", ".env", "env file path")
}

func main() {
	t := time.Now()
	fmt.Printf("[%s] app is starting ...\n", t.Format("2006-01-02 15:04:05"))

	debug.Start()
	debug.StartMemProfile()

	flag.Parse()
	if err := app.LoadConfig(); err != nil {
		fmt.Printf("failed to load config: %v\n", err)
		fmt.Println("using default config file")
	}

	database, err := db.NewPostgres(app.AppConfig.DB)
	if err != nil {
		log.Fatal(err)
	}

	h := ws.NewHub()
	go h.Run()

	r := router.New(app.AppConfig.Origins, h)
	//wsHub := ws.NewHub()

	application := app.CreateApp(
		app.WithTimeoutConfig(app.AppConfig.HTTP),
		app.WithDB(database),
		app.WithRouter(r.Handler()),
		//app.WithWebSocket(wsHub),
		app.WithJwtGenerator(app.AppConfig.JWT.Secret),
		app.WithAuthVerifier(app.AppConfig.JWT.Secret),
	)

	go func() {
		fmt.Printf("starting app on %s\n", app.AppConfig.Addr)
		if err := application.Start(app.AppConfig.Addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v", err)
		}
	}()

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	log.Println("Shutdown signal received, shutting down...")

	h.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	application.Stop(ctx)

	debug.StopMemProfile()
	fmt.Printf("[%s] app is stopped, bye!\n", t.Format("2006-01-02 15:04:05"))
}
