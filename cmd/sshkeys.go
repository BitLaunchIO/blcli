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

	"github.com/spf13/cobra"
)

// SSHKey sets up the ssh key command and subcommands
func SSHKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sshkey",
		Short:   "Manage SSH Keys",
		Long:    `Use the subcommands to list, create, or delete ssh keys.`,
		Aliases: []string{"k"},
	}

	cmd.AddCommand(sshKeyList)
	cmd.AddCommand(sshKeyDelete)
	cmd.AddCommand(sshKeyCreate)

	sshKeyCreate.Flags().StringP("name", "n", "", "name for the new key")
	sshKeyCreate.Flags().StringP("content", "c", "", "ssh key content")
	sshKeyCreate.MarkFlagRequired("name")
	sshKeyCreate.MarkFlagRequired("content")

	return cmd
}

var sshKeyList = &cobra.Command{
	Use:     "list",
	Short:   "List ssh keys on your account",
	Long:    ``,
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		servers, err := client.SSHKey.List()
		if err != nil {
			fmt.Printf("Error listing ssh keys : %v\n", err)
			os.Exit(1)
		}

		printer.Output(servers)
	},
}

var sshKeyDelete = &cobra.Command{
	Use:     "delete",
	Short:   "Permanently delete an ssh key",
	Long:    `delete <key-id>`,
	Aliases: []string{"delete", "d", "del", "rm"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide am ssh key ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		err := client.SSHKey.Delete(id)
		if err != nil {
			fmt.Printf("Error deleting ssh key : %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Deleted ssh key")
	},
}

var sshKeyCreate = &cobra.Command{
	Use:     "create",
	Short:   "Create a new ssh key",
	Long:    ``,
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		opts := gobitlaunch.SSHKey{}
		opts.Name, _ = cmd.Flags().GetString("name")
		opts.Content, _ = cmd.Flags().GetString("content")

		key, err := client.SSHKey.Create(&opts)
		if err != nil {
			fmt.Printf("Error creating ssh key : %v\n", err)
			os.Exit(1)
		}

		printer.Output(key)
	},
}
