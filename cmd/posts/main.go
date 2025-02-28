package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"

	"github.com/zhenyanesterkova/nepblog/internal/app/config"
	"github.com/zhenyanesterkova/nepblog/internal/app/logger"
	"github.com/zhenyanesterkova/nepblog/internal/app/storage"
	"github.com/zhenyanesterkova/nepblog/internal/http/handlers"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func run() error {
	cfg := config.New()
	cfg.Build()

	loggerInst := logger.NewLogrusLogger()
	err := loggerInst.SetLevelForLog(cfg.LConfig.Level)
	if err != nil {
		loggerInst.LogrusLog.Errorf("can not parse log level: %v", err)
		return fmt.Errorf("parse log level error: %w", err)
	}

	store, err := storage.NewStore(cfg, loggerInst)
	if err != nil {
		loggerInst.LogrusLog.Errorf("failed create storage: %v", err)
		return fmt.Errorf("failed create storage: %w", err)
	}

	defer func() {
		err := store.Close()
		if err != nil {
			loggerInst.LogrusLog.Errorf("can not close storage: %v", err)
		}
	}()

	router := chi.NewRouter()

	repoHandler := handlers.NewRepositorieHandler(store, loggerInst)
	repoHandler.InitChiRouter(router)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	loggerInst.LogrusLog.Infof("Start Server on %s", cfg.SConfig.Address)
	go func() {
		if err := http.ListenAndServe(cfg.SConfig.Address, router); err != nil {
			loggerInst.LogrusLog.Errorf("server error: %v", err)
		}
	}()

	s := <-c
	loggerInst.LogrusLog.Info("Got signal: ", s)

	return nil
}
