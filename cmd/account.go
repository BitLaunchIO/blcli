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
	return &cobra.Command{
		Use:   "account",
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
}
