/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a link from the table",
	Long:  `Deletes a record from the RedirectsLink table`,
	Run: func(cmd *cobra.Command, args []string) {
		introDelMessages(FromPath)

		client := &http.Client{}

		reqURL := BaseURL + FromPath

		req, err := http.NewRequest(http.MethodDelete, reqURL, nil)
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json; charset=utf8")
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	delCmd.Flags().StringVarP(&FromPath, "from", "f", "", "The path on your domain to link from")
	delCmd.MarkFlagRequired("from")
}

func introDelMessages(f string) {
	fmt.Println("Deleting the following to the link shortener table")
	fmt.Printf("\t- From Path: %v\n", f)
}
