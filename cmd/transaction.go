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

	"./printer"
	"github.com/bitlaunchio/gobitlaunch"
	"github.com/mdp/qrterminal"
	"github.com/spf13/cobra"
)

// Transaction sets up the server command and subcommands
func Transaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "transaction",
		Short:   "Manage transactions",
		Long:    `Use the subcommands to get, list, or create transactions.`,
		Aliases: []string{"t"},
	}

	cmd.AddCommand(transactionCreate)
	cmd.AddCommand(transactionList)
	cmd.AddCommand(transactionGet)
	cmd.AddCommand(transactionQRCode)
	transactionGet.Flags().Bool("qr", false, "output transaction as qr code to terminal")
	transactionCreate.Flags().Bool("qr", false, "output transaction as qr code to terminal")
	transactionCreate.Flags().BoolP("lightning", "l", false, "optionally use lightning network valid for BTC and LTC up to 0.042 BTC or equivalent.")
	transactionList.Flags().IntP("page", "p", 1, "page number")
	transactionList.Flags().IntP("items", "i", 25, "number of items per page")

	return cmd
}

var transactionCreate = &cobra.Command{
	Use:     "create",
	Short:   "Create a new transaction",
	Long:    `create <amount-usd> <BTC, LTC, ETH, BCH, NANO, TRX, SRN, TEL>`,
	Aliases: []string{"c", "create"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("create <amount-usd> <BTC, LTC, ETH, BCH, NANO, TRX, SRN, TEL>")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		usd := args[0]
		symbol := args[1]
		usdInt, err := strconv.Atoi(usd)
		if err != nil {
			fmt.Println("Please specify USD as an integer")
			os.Exit(1)
		}
		ln, _ := cmd.Flags().GetBool("lightning")
		if ln && (symbol != "BTC" && symbol != "LTC") {
			fmt.Println("Lightning network only available for BTC and LTC")
			os.Exit(1)
		}
		transaction, err := client.Transaction.Create(&gobitlaunch.CreateTransactionOptions{
			AmountUSD:        usdInt,
			CryptoSymbol:     symbol,
			LightningNetwork: ln,
		})
		if err != nil {
			fmt.Printf("Error creating a new transaction : %v\n", err)
			os.Exit(1)
		}

		qr, _ := cmd.Flags().GetBool("qr")
		if qr {
			if len(transaction.Address) == 0 || len(transaction.AmountCrypto) == 0 {
				fmt.Println("Unable to generate a QR Code for this type of transaction.")
				os.Exit(1)
			}
			s := fmt.Sprintf("bitcoin:%s?amount=%s", transaction.Address, transaction.AmountCrypto)
			qrterminal.Generate(s, qrterminal.L, os.Stdout)
			return
		}

		printer.Output(transaction)
	},
}

var transactionGet = &cobra.Command{
	Use:     "get",
	Short:   "Get information for a single transaction",
	Long:    `get <transaction-id>`,
	Aliases: []string{"g", "show"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a transaction ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		transaction, err := client.Transaction.Show(id)
		if err != nil {
			fmt.Printf("Error getting transaction : %v\n", err)
			os.Exit(1)
		}

		qr, _ := cmd.Flags().GetBool("qr")
		if qr {
			if len(transaction.Address) == 0 || len(transaction.AmountCrypto) == 0 {
				fmt.Println("Unable to generate a QR Code for this type of transaction.")
				os.Exit(1)
			}
			s := fmt.Sprintf("bitcoin:%s?amount=%s", transaction.Address, transaction.AmountCrypto)
			qrterminal.Generate(s, qrterminal.L, os.Stdout)
			return
		}

		printer.Output(transaction)
	},
}

var transactionList = &cobra.Command{
	Use:     "list",
	Short:   "List transactions on your account",
	Long:    `list --page [page-number|1] --items [items-per-page|25]`,
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		page, _ := cmd.Flags().GetInt("page")
		items, _ := cmd.Flags().GetInt("items")
		transactions, err := client.Transaction.List(page, items)
		if err != nil {
			fmt.Printf("Error listing transactions : %v\n", err)
			os.Exit(1)
		}

		printer.Output(transactions)
	},
}

var transactionQRCode = &cobra.Command{
	Use:   "qr",
	Short: "Generate a QR code for a transaction",
	Long:  `qr <transaction-id>`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a transaction ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		transaction, err := client.Transaction.Show(id)
		if err != nil {
			fmt.Printf("Error getting transaction : %v\n", err)
			os.Exit(1)
		}

		if len(transaction.Address) == 0 || len(transaction.AmountCrypto) == 0 {
			fmt.Println("Unable to generate a QR Code for this type of transaction.")
			os.Exit(1)
		}

		s := fmt.Sprintf("bitcoin:%s?amount=%s", transaction.Address, transaction.AmountCrypto)

		qrterminal.Generate(s, qrterminal.L, os.Stdout)
	},
}
