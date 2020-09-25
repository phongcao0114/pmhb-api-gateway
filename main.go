package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pmhb-api-gateway/internal/api"
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/pkg/klog"

	"strings"
	"syscall"
	"time"
)

// Main function
func main() {
	// 1. Safely close all log file writers (if exists)
	defer klog.Close()

	configPath := flag.String("config", "configs", "set configs path, default as: 'configs'")
	state := flag.String("state", "dev", "set working environment")
	port := flag.String("port", "10000", "port number")
	flag.Parse()

	// 2. Allow override state of the app via environment variable
	appState := os.Getenv("APP_STATE")
	if len(strings.TrimSpace(appState)) > 0 {
		*state = appState
	}

	// 3. Prepare logger
	KLogger := klog.WithPrefix("main")
	KLogger.WithFields(map[string]interface{}{
		"state":  *state,
		"port":   *port,
		"config": *configPath,
	}).Info("starting server")

	// 4. Get new configuration
	cfg, err := config.New(*configPath, *state)
	if err != nil {
		KLogger.Panicf("[main] Failed to load config, err: %v", err)
	}
	config.Config = cfg

	// 5. Set singleton
	utils.ResponseAppID = config.Config.AppID
	config.InitRandomProfileUserID()

	// 6. Load location
	bkkLocation, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		KLogger.Panicf("[main] Failed to load location, err: %v", err)
	}
	utils.BKKLocation = bkkLocation

	srvCtx, srvCancel := context.WithCancel(context.Background())

	// 7. Start http server
	router, err := api.NewRouter(config.Config)
	if err != nil {
		KLogger.Errorf("failed to init router, err: %v", err)
	}
	srv := &http.Server{
		Addr:              fmt.Sprint(":", *port),
		Handler:           router,
		ReadTimeout:       config.Config.HTTPServer.ReadTimeout,
		WriteTimeout:      config.Config.HTTPServer.WriteTimeout,
		ReadHeaderTimeout: config.Config.HTTPServer.ReadHeaderTimeout,
	}

	// 8. Listen HTTP request on background
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			KLogger.Errorf("listen: %s\n", err)
		}
	}()

	// 9. Graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	<-signals
	KLogger.Info("shutting down http server...")
	if err := srv.Shutdown(srvCtx); err != nil {
		KLogger.Panicln("http server shutdown with error:", err)
	}
	srvCancel()
}
