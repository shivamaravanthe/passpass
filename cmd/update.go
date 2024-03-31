/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/shivamaravanthe/passpass/file"
	"github.com/spf13/cobra"
)

var (
	newName string
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the password",
	Long:  "Updates the password if the key is correct",
	Run: func(cmd *cobra.Command, args []string) {
		key, success := promptPassword()
		if !success {
			fmt.Println("Failed to get Key")
			fmt.Println("Save Failed")
			return
		}
		data, success := promptData()
		if !success {
			fmt.Println("Failed to get Data")
			fmt.Println("Save Failed")
			return
		}

		file.Update(Name, newName, data, key, Store)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&Name, "name", "o", "", "Name of password.(Name will not be encrypted)")
	if err := updateCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	updateCmd.Flags().StringVarP(&newName, "newName", "n", "", "New Name of password.(Name will not be encrypted)")
}
