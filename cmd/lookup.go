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
	"strings"

	"github.com/paulfarver/dub/pkg/behindthename"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	exact bool
)

// lookupCmd represents the lookup command
var lookupCmd = &cobra.Command{
	Use:   "lookup [name]",
	Short: "This will return information about a given name",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := behindthename.NewClient(viper.GetString("api.token"), http.DefaultClient)
		b, err := client.Lookup(context.Background(), args[0], behindthename.LookupParams{
			Exact: exact,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		printLookups(*b)
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)

	lookupCmd.Flags().BoolVarP(&exact, "exact", "e", false, "whether the name supplied is exact (meaning there are no missing diacritics)")
}

func printLookups(l []behindthename.LookupResponseElement) {
	switch viper.GetString(confOutputFormat) {
	case "json":
		d, _ := json.Marshal(l)
		fmt.Println(string(d))
	case "yaml":
		d, _ := yaml.Marshal(l)
		fmt.Println(string(d))
	case "text":
		fallthrough
	default:
		for _, v := range l {
			printLookup(v)
		}
	}
}

func printLookup(v behindthename.LookupResponseElement) {
	usageStrings := []string{}
	for _, v := range v.Usages {
		usageStrings = append(usageStrings, fmt.Sprintf("%s %s", v.UsageFull, v.UsageGender))
	}
	fmt.Printf("%s %s [%s]\n", v.Name, v.Gender, strings.Join(usageStrings, ", "))
}
