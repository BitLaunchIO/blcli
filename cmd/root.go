/*
Copyright 2020 The blcli Authors All rights reserved.
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
	"errors"
	"fmt"
	"os"

	"./printer"
	"github.com/bitlaunchio/gobitlaunch"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	client  *gobitlaunch.Client
	cfgFile string
	token   string
	format  string

	rootCmd = &cobra.Command{
		Use:   "blcli",
		Short: "blcli is a command-line interface for BitLaunch.io",
		Long:  ``,
	}
)

func hostID(name string) (int, error) {
	var err error
	var h int
	switch name {
	case "bitlaunch", "bl":
		h = 4
	case "digitalocean", "do":
		h = 0
	case "vultr", "v":
		h = 1
	case "linode", "l":
		h = 2
	default:
		err = errors.New("invalid host")
	}
	return h, err
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initClient)
	cobra.OnInitialize(initPrinter)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blcli.yaml)")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "API authentication token")
	//rootCmd.PersistentFlags().StringVar(&format, "format", "json", "output format. can be: kv, csv or json (default)")
	rootCmd.MarkFlagRequired("token")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(Account())
	rootCmd.AddCommand(Server())
	rootCmd.AddCommand(Transaction())
	rootCmd.AddCommand(CreateOptions())
	rootCmd.AddCommand(SSHKey())
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		// Search config in home directory with name ".blcli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".blcli")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initClient() {
	if versionCmd.CalledAs() == "version" {
		return
	}
	if len(token) == 0 {
		token = os.Getenv("BL_API_TOKEN")
		if len(token) == 0 {
			fmt.Println("You must specify your API token with either the --token parameter or by exporting it as an environment variable:")
			fmt.Println("export BL_API_TOKEN='<your_token_here>'")
			os.Exit(1)
		}
	}

	client = gobitlaunch.NewClient(token)
}

func initPrinter() {
	printer.Format = "json"
}
