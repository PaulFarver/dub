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
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	confOutputFormat = "output.format"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dub",
	Short: "A quick cli to call the behindthename api",
	Long: ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	viper.SetDefault(confOutputFormat, "text")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dub.yaml)")
	rootCmd.PersistentFlags().StringP("output", "o", viper.GetString(confOutputFormat), "Output format")

	viper.BindPFlag(confOutputFormat, rootCmd.PersistentFlags().Lookup("output"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Find current directory
		wd, err := os.Getwd()
		cobra.CheckErr(err)
		viper.AddConfigPath(wd)

		// Search config in home directory with name ".dub" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".dub")
	}

	viper.AutomaticEnv() // read in environment variables that match

	viper.ReadInConfig()
}

type PrinatableResponse interface {
	GetNames() []string
}

func printNames(resp PrinatableResponse) {
	switch viper.GetString(confOutputFormat) {
	case "json":
		d, _ := json.Marshal(resp)
		fmt.Println(string(d))
	case "yaml":
		d, _ := yaml.Marshal(resp)
		fmt.Println(string(d))
	case "text":
		fallthrough
	default:
		for _, n := range resp.GetNames() {
			fmt.Println(n)
		}
	}
}
