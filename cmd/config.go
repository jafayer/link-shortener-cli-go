/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jafayer/shorten/pkg/config"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var newLink string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Check or set the URL of your CLI methods cloud function",
	Long: `Check or set the URL of your CLI methods cloud function.
	
	Run without the --link flag to get the current link value.
	
	Supply the link flag to set a new link`,
	Run: func(cmd *cobra.Command, args []string) {
		rootLink := config.GetRootLink()
		if rootLink != newLink && newLink != "" {
			fmt.Printf("Setting new CLI methods link to %v\n", newLink)
			viper.Set("LINK", newLink)
			err := viper.WriteConfig()
			if err != nil {
				panic(err)
			}
		} else {
			if !config.RootLinkIsSet() {
				fmt.Println("\nCLI methods link is not set! use `shorten config --link <link>` to set it")
			} else {
				fmt.Printf("The current CLI methods link is: %v\n", rootLink)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	configCmd.Flags().StringVarP(&newLink, "link", "s", "", "Set the lambda root link")
}
