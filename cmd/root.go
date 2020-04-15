package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "django-dep",
	Short: "Runs a daemon to manage django based server",
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}

// Run function lets you run the commands
func Run(args []string) error {
	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}
