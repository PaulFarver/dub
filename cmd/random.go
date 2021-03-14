/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/paulfarver/dub/pkg/behindthename"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	gender        string
	usage         string
	number        int
	randomSurname bool
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get random names",
	Long:  `Get a list of random names`,
	Run: func(cmd *cobra.Command, args []string) {
		client := behindthename.NewClient(viper.GetString("api.token"), http.DefaultClient)
		resp, err := client.RandomName(context.TODO(), behindthename.RandomNameParameters{
			Gender:        gender,
			Usage:         usage,
			Number:        number,
			RandomSurname: randomSurname,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		printNames(resp)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	randomCmd.Flags().StringVar(&gender, "gender", "", "restrict names to a specific gender")
	randomCmd.Flags().StringVar(&usage, "usage", "", "restrict names to a specific usage such as sla for slavic")
	randomCmd.Flags().IntVar(&number, "number", 2, "amount of names to get")
	randomCmd.Flags().BoolVar(&randomSurname, "surname", false, "generate surnames")
}

func printNames(resp *behindthename.RandomNameResponse) {
	switch viper.GetString("output.format") {
	case "json":
		d, _ := json.Marshal(resp)
		fmt.Println(string(d))
	case "yaml":
		d, _ := yaml.Marshal(resp)
		fmt.Println(string(d))
	case "text":
		fallthrough
	default:
		for _, n := range resp.Names {
			fmt.Println(n)
		}
	}
}
