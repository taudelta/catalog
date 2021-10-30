package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
	"github.com/stanyx/catalog/app/handler"
	"github.com/stanyx/catalog/config"
	"github.com/stanyx/catalog/internal/storage"
)

func main() {

	var cfg zap.Config

	log.SetFlags(log.Llongfile | log.LstdFlags)

	appConfig, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	loggerCfg, err := json.MarshalIndent(appConfig.Logger, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(loggerCfg))

	if err := json.Unmarshal(loggerCfg, &cfg); err != nil {
		log.Fatal(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	logger.Debug("Start")

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	r := mux.NewRouter()

	handler.RegisterEndpoints(r)

	if err := storage.InitDB(storage.DatabaseOptions{
		Host:   appConfig.Database.Host,
		Port:   appConfig.Database.Port,
		User:   appConfig.Database.User,
		Pass:   appConfig.Database.Pass,
		Dbname: appConfig.Database.Dbname,
	}); err != nil {
		logger.Sugar().Fatalf("database error: %s", err)
	}

	server := http.Server{
		Addr:    ":10000",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Sugar().Fatalf("server error: %s", err)
		}
	}()

	select {
	case sig := <-sigCh:

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Sugar().Errorf("server shutdown error %s", sig)
		}
		logger.Sugar().Infof("stopped by signal %s", sig)
		return
	}

}
