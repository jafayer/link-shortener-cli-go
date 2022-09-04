/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type toURL struct {
	Value string `json:"S"`
}

type redirectsTable struct {
	ToURL toURL `json:"toURL"`
}

type response struct {
	RedirectsTable []redirectsTable `json:"RedirectsTable"`
}

type responseBody struct {
	Response response `json:"Responses"`
}

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Link string = viper.GetString("LINK")
		var LinkNotSet bool = Link == ""

		if LinkNotSet {
			fmt.Println("Error: Cannot execute because link is not set")
			fmt.Println("Please use `shorten config --link` to add a CLI shortener remote link")
			return
		}

		introLsMessages(FromPath)

		client := &http.Client{}

		reqURL := Link + FromPath

		req, err := http.NewRequest(http.MethodGet, reqURL, nil)
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json; charset=utf8")
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		jsonResponse, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var jsonData responseBody

		err = json.Unmarshal([]byte(jsonResponse), &jsonData)
		if err != nil {
			panic(err)
		}

		if len(jsonData.Response.RedirectsTable) < 1 {
			fmt.Println("There was no link with that fromPath!")
			return
		}
		fmt.Printf("\t- To URL:\t%v", jsonData.Response.RedirectsTable[0].ToURL.Value)

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	lsCmd.Flags().StringVarP(&FromPath, "from", "f", "", "The path on your domain to link from")
	lsCmd.MarkFlagRequired("from")
}

func introLsMessages(f string) {
	fmt.Println("Retrieving from the server:")
	fmt.Printf("\t- From Path:\t%v\n", f)
}
