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
	"strconv"
	"strings"

	"./printer"
	"github.com/bitlaunchio/gobitlaunch"

	"github.com/spf13/cobra"
)

// Server sets up the server command and subcommands
func Server() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server",
		Short:   "Manage your virtual machines",
		Long:    `Use the subcommands to get, list, create, or destroy servers.`,
		Aliases: []string{"s"},
	}

	cmd.AddCommand(serverGet)
	cmd.AddCommand(serverList)
	cmd.AddCommand(serverDestroy)
	cmd.AddCommand(serverCreate)
	cmd.AddCommand(serverRebuild)
	cmd.AddCommand(serverResize)
	cmd.AddCommand(serverRestart)
	cmd.AddCommand(serverProtection)
	cmd.AddCommand(serverSetPorts)

	serverCreate.Flags().StringP("name", "n", "", "name for the new server")
	serverCreate.Flags().StringP("host", "t", "", "target provider/host name: bitlaunch, digitalocean, vultr or linode")
	serverCreate.Flags().StringP("image", "i", "", "image/app id")
	serverCreate.Flags().StringP("size", "s", "", "plan/size id")
	serverCreate.Flags().StringP("region", "r", "", "region id")
	serverCreate.Flags().StringSliceP("sshkey", "k", []string{}, "ssh key ids, comma separated for more than one")
	serverCreate.Flags().StringP("password", "p", "", "password")

	serverRebuild.Flags().StringP("image", "i", "", "image/app id")
	serverRebuild.Flags().StringP("description", "d", "", "image/app description")

	serverResize.Flags().StringP("size", "s", "", "plan/size id")

	serverSetPorts.Flags().StringP("ports", "p", "", "port:protocol, comma separated for more than one")

	serverCreate.MarkFlagRequired("name")
	serverCreate.MarkFlagRequired("host")
	serverCreate.MarkFlagRequired("image")
	serverCreate.MarkFlagRequired("size")
	serverCreate.MarkFlagRequired("region")

	serverRebuild.MarkFlagRequired("image")
	serverRebuild.MarkFlagRequired("description")

	serverResize.MarkFlagRequired("size")

	serverSetPorts.MarkFlagRequired("ports")

	return cmd
}

var serverGet = &cobra.Command{
	Use:     "get",
	Short:   "Get information for a single server",
	Long:    `get <server-id>`,
	Aliases: []string{"g", "show"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a server ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		server, err := client.Server.Show(id)
		if err != nil {
			fmt.Printf("Error getting server : %v\n", err)
			os.Exit(1)
		}

		printer.Output(server)
	},
}

var serverList = &cobra.Command{
	Use:     "list",
	Short:   "List servers on your account",
	Long:    ``,
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		servers, err := client.Server.List()
		if err != nil {
			fmt.Printf("Error listing servers : %v\n", err)
			os.Exit(1)
		}

		printer.Output(servers)
	},
}

var serverDestroy = &cobra.Command{
	Use:     "destroy",
	Short:   "Permanently delete a server",
	Long:    `destroy <server-id>`,
	Aliases: []string{"delete", "d", "del", "rm"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a server ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		err := client.Server.Destroy(id)
		if err != nil {
			fmt.Printf("Error destroying server : %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Deleted server")
	},
}

var serverCreate = &cobra.Command{
	Use:     "create",
	Short:   "Create a new server",
	Long:    ``,
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		opts := gobitlaunch.CreateServerOptions{}
		opts.Name, _ = cmd.Flags().GetString("name")
		host, _ := cmd.Flags().GetString("host")
		opts.HostImageID, _ = cmd.Flags().GetString("image")
		opts.SizeID, _ = cmd.Flags().GetString("size")
		opts.RegionID, _ = cmd.Flags().GetString("region")
		opts.SSHKeys, _ = cmd.Flags().GetStringSlice("sshkey")
		opts.Password, _ = cmd.Flags().GetString("password")
		opts.InitScript, _ = cmd.Flags().GetString("initscript")

		// validate
		if len(opts.Password) == 0 && len(opts.SSHKeys) == 0 {
			fmt.Println("You must provide either --sshkey or --password")
			os.Exit(1)
		}

		var err error
		opts.HostID, err = hostID(host)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		server, err := client.Server.Create(&opts)
		if err != nil {
			fmt.Printf("Error creating server : %v\n", err)
			os.Exit(1)
		}

		printer.Output(server)
	},
}

var serverRebuild = &cobra.Command{
	Use:     "rebuild",
	Short:   "Rebuild a server",
	Long:    `rebuild <server-id>`,
	Aliases: []string{},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a server ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		opts := gobitlaunch.RebuildOptions{}
		opts.ID, _ = cmd.Flags().GetString("image")
		opts.Description, _ = cmd.Flags().GetString("description")

		err := client.Server.Rebuild(id, &opts)
		if err != nil {
			fmt.Printf("Error rebuilding server : %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Rebuilding server")
	},
}

var serverResize = &cobra.Command{
	Use:     "resize",
	Short:   "Resize a server",
	Long:    `resize <server-id>`,
	Aliases: []string{},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a server ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		sizeID, _ := cmd.Flags().GetString("size")

		err := client.Server.Resize(id, sizeID)
		if err != nil {
			fmt.Printf("Error resizing server : %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Resizing server")
	},
}

var serverRestart = &cobra.Command{
	Use:     "restart",
	Short:   "Restart a server",
	Long:    `restart <server-id>`,
	Aliases: []string{"reboot"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a server ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		err := client.Server.Restart(id)
		if err != nil {
			fmt.Printf("Error restarting server : %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Restarted server")
	},
}

var serverProtection = &cobra.Command{
	Use:     "protection",
	Short:   "Protect a server",
	Long:    `protection <server-id> [enable true e] or [disable false d]`,
	Aliases: []string{"protect"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a server ID")
		}
		if len(args) < 2 {
			return errors.New("please provide a protection state")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		server, err := client.Server.Protection(id, func() bool {
			if args[1] == "enable" || args[1] == "true" || args[1] == "e" {
				return true
			} else if args[1] == "disable" || args[1] == "false" || args[1] == "d" {
				return false
			}

			fmt.Println("Invalid protection state")
			os.Exit(1)
			return false
		}())
		if err != nil {
			fmt.Printf("Error resizing server : %v\n", err)
			os.Exit(1)
		}

		printer.Output(server)
	},
}

var serverSetPorts = &cobra.Command{
	Use:     "setports",
	Short:   "Set ports for a protected server",
	Long:    `setports <server-id>`,
	Aliases: []string{"ports"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a server ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		ports, _ := cmd.Flags().GetString("ports")

		portItems := strings.Split(ports, ",")
		portList := []gobitlaunch.Ports{}

		for _, port := range portItems {
			portObj := strings.Split(port, ":")

			num, err := strconv.Atoi(portObj[0])
			if err != nil {
				fmt.Printf("Error setting server ports : %v\n", err)
				os.Exit(1)
			}

			portList = append(portList, gobitlaunch.Ports{
				PortNumber: num,
				Protocol:   portObj[1],
			})
		}

		server, err := client.Server.SetPorts(id, &portList)
		if err != nil {
			fmt.Printf("Error setting server ports : %v\n", err)
			os.Exit(1)
		}

		printer.Output(server)
	},
}
