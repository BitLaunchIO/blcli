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
	"fmt"
	"os"

	"./printer"

	"github.com/spf13/cobra"
)

// Account sets up the account command
func Account() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "Retrieve account information",
		Long:  `Use the subcommands to display information about your account.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				account, err := client.Account.Show()
				if err != nil {
					fmt.Printf("Error getting account information : %v", err)
					os.Exit(1)
				}

				printer.Output(account)
			}
		},
	}

	cmd.AddCommand(accountShow)
	cmd.AddCommand(accountUsage)
	cmd.AddCommand(accountHistory)

	accountUsage.Flags().StringP("period", "p", "latest", "filter for period, format: YYYY-MM or latest")

	accountHistory.Flags().IntP("page", "p", 1, "page number of history results to show")
	accountHistory.Flags().IntP("items", "i", 25, "how many history results to show")

	return cmd
}

var accountShow = &cobra.Command{
	Use:   "show",
	Short: "Retrieve account information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		account, err := client.Account.Show()
		if err != nil {
			fmt.Printf("Error getting account information : %v", err)
			os.Exit(1)
		}

		printer.Output(account)
	},
}

var accountUsage = &cobra.Command{
	Use:   "usage",
	Short: "Retrieve account usage information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		period, _ := cmd.Flags().GetString("period")

		usage, err := client.Account.Usage(period)

		if err != nil {
			fmt.Printf("Error getting account usage information : %v", err)
			os.Exit(1)
		}

		printer.Output(usage)
	},
}

var accountHistory = &cobra.Command{
	Use:   "history",
	Short: "Retrieve account history information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		page, _ := cmd.Flags().GetInt("page")
		items, _ := cmd.Flags().GetInt("items")

		history, err := client.Account.History(page, items)
		if err != nil {
			fmt.Printf("Error getting account history information : %v", err)
			os.Exit(1)
		}

		printer.Output(history)
	},
}
