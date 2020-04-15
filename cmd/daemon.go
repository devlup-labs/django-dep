package cmd

import (
	"context"
	"fmt"
	"github.com/devlup-labs/django-dep/config"
	"github.com/devlup-labs/django-dep/logger"
	"github.com/devlup-labs/django-dep/router"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start the django-dep daemon",
	Run:   daemonCmdF,
}

func daemonCmdF(command *cobra.Command, args []string) {
	logger.Infof("Starting daemon on port %d", config.Port())
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port()),
		Handler: router.NewRouter(),
	}

	go func() {
		err := s.ListenAndServe()
		if err != http.ErrServerClosed && err != nil {
			logger.Errorf("error while starting daemon: %s", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)

	<-sig
	logger.Info("daemon shutting down")

	err := s.Shutdown(context.Background())
	if err != nil {
		logger.Error(err.Error())
	}
	close(sig)
	logger.Info("daemon shutdown complete")
}
