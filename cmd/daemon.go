package cmd

import (
	"fmt"
	"github.com/devlup-labs/django-dep/config"
	"github.com/devlup-labs/django-dep/router"
	"github.com/spf13/cobra"
	"net/http"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start the django-dep daemon",
	Run:   daemonCmdF,
}

func daemonCmdF(command *cobra.Command, args []string) {
	r := router.NewRouter()
	_ = http.ListenAndServe(fmt.Sprintf(":%d", config.Port()), r)
}
