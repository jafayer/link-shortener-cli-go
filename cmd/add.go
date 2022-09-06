/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jafayer/shorten/pkg/config"
	"github.com/spf13/cobra"
)

type Metadata struct {
	HTTPStatusCode  uint8  `json:"httpStatusCode"`
	RequestId       string `json:"requestId"`
	Attempts        uint8  `json:"attempts"`
	TotalRetryDelay uint8  `json:"totalRetryDelay"`
}

type BodyStruct struct {
	Metadata Metadata `json:"$metadata"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a link to the shortener",
	Long: `Call "add" to PUT a link in the shortener.
	Supply a from path and a to URL as below:

		shorten add -f path -t https://example.com
	`,
	Run: func(cmd *cobra.Command, args []string) {

		config.ErrRootLinkNotSet()
		if !config.RootLinkIsSet() {
			return
		}

		introAddingMessages(FromPath, ToURL)

		data := map[string]string{
			"fromPath":   FromPath,
			"toURL":      ToURL,
			"httpMethod": "PUT",
		}

		body := map[string]interface{}{
			"body":       data,
			"httpMethod": "PUT",
		}

		// intialize client
		client := &http.Client{}

		json_data, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest(http.MethodPut, config.GetRootLink(), bytes.NewBuffer(json_data))
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json; charset=utf8")
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		jsonBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var bodyMap map[string]string

		err = json.Unmarshal([]byte(jsonBody), &bodyMap)
		if err != nil {
			panic(err)
		}

		metaString := bodyMap["body"]

		var respData BodyStruct
		err = json.Unmarshal([]byte(metaString), &respData)
		if err != nil {
			panic(err)
		}

		isSuccess := respData.Metadata.HTTPStatusCode == 200
		hasAttempts := respData.Metadata.Attempts > 0

		if isSuccess && hasAttempts {
			fmt.Println("Successfully added!")
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringVarP(&FromPath, "from", "f", "", "The path on your domain to link from")
	addCmd.Flags().StringVarP(&ToURL, "to", "t", "", "The full URL to link to")
	addCmd.MarkFlagRequired("from")
	addCmd.MarkFlagRequired("to")
	addCmd.MarkFlagsRequiredTogether("from", "to")
}

func introAddingMessages(f string, t string) {
	fmt.Println("Adding the following to the link shortener table")
	fmt.Printf("\t- From Path: %v\n", f)
	fmt.Printf("\t- To URL: %v\n", t)
	fmt.Println("Please wait...")
}
