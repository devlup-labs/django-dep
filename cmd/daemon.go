package cmd

import (
	"github.com/devlup-labs/django-dep/handler"
	"github.com/spf13/cobra"
	"net/http"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start the django-dep daemon",
	Run:   daemonCmdF,
}

func daemonCmdF(command *cobra.Command, args []string) {
	http.HandleFunc("/", handler.Deploy)
	http.HandleFunc("/ping", handler.Ping)
	http.ListenAndServe(":8080", nil)
}
