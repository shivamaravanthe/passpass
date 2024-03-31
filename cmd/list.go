/*
Copyright © 2024 Shivanand Shenoy <vshivanand2@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/shivamaravanthe/passpass/file"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the saved passwords",
	Long:  "List all the saved passwords",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("-----Names of Saved Passwords------")
		file.List(Store)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
