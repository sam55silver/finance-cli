/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/sam55silver/finance-cli/lib"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: call,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func call(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatalln("No noun. usage: category, transactions")
	}

	db := lib.DBConnect("./finance.db")
	defer db.Close()

	noun := args[0]

	if noun == "category" {
		for i := 1; i < len(args); i++ {
			id := db.CreateCategory(args[i])
			fmt.Println("Created Food category, id:", id)
		}
	} else if noun == "transactions" {
		var input string
		for {
			fmt.Print("Enter transaction: ")
			fmt.Scan(&input)

			if input == "stop" {
				break
			} else {
				fmt.Println("You entered,", input)
			}
		}
	}
}
