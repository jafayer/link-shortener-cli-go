/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const BaseURL string = "https://ao4qyu2pr0.execute-api.us-east-1.amazonaws.com/default/LinkShortenerCLIMethods-LinkShortenerCLIMethods-TiijGWZnpMZ5/"

var FromPath string
var ToURL string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shorten",
	Short: "A CLI utility for link shortening",
	Long: `A CLI for interacting with a Link Shortening API.
	
	Perform CRUD operations with the add, del, and ls utilities.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.link-shortener-cli-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
