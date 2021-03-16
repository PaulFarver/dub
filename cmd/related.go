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
	"fmt"
	"net/http"
	"os"

	"github.com/paulfarver/dub/pkg/behindthename"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	relatedGender string
	relatedUsage  string
)

// relatedCmd represents the related command
var relatedCmd = &cobra.Command{
	Use:   "related",
	Short: "This will return potential aliases for a given name",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := behindthename.NewClient(viper.GetString("api.token"), http.DefaultClient)
		resp, err := client.RelatedNames(context.TODO(), args[0], behindthename.RelatedNamesParameters{
			Gender: relatedGender,
			Usage:  relatedUsage,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		printNames(resp)
	},
}

func init() {
	rootCmd.AddCommand(relatedCmd)

	relatedCmd.Flags().StringVarP(&relatedGender, "gender", "g", "", "restrict names to a specific gender")
	relatedCmd.Flags().StringVarP(&relatedUsage, "usage", "u", "", "restrict names to a specific usage such as eng for english")
}
