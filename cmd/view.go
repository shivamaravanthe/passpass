/*
Copyright © 2024 Shivanand Shenoy <vshivanand2@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/shivamaravanthe/passpass/file"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Displays the Password",
	Long:  "If given Key is correct displays the password",
	Run: func(cmd *cobra.Command, args []string) {
		key, success := promptPassword()
		if success {
			file.View(Name, key, Store)
		} else {
			fmt.Println("Save Failed")
		}

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	viewCmd.Flags().StringVarP(&Name, "name", "n", "", "Name of password.(Name will not be encrypted)")
	if err := viewCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
}
