/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var link string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Check or set the URL of your CLI methods cloud function",
	Long: `Check or set the URL of your CLI methods cloud function.
	
	Run without the --link flag to get the current link value.
	
	Supply the link flag to set a new link`,
	Run: func(cmd *cobra.Command, args []string) {
		if link != "" {
			fmt.Printf("Setting new CLI methods link to %v\n", link)
			viper.Set("LINK", link)
			err := viper.WriteConfig()
			if err != nil {
				panic(err)
			}
		} else {

			fmt.Println("The current CLI methods link is:")
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
	configCmd.Flags().StringVarP(&link, "link", "s", "", "Set the lambda root link")
}
