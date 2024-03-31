/*
Copyright © 2024 Shivanand Shenoy <vshivanand2@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/shivamaravanthe/passpass/file"
	"github.com/spf13/cobra"
)

var (
	Name string
)

// saveCmd represents the save command
var SaveCmd = &cobra.Command{
	Use:   "save",
	Short: "Saves the password",
	Long:  "Saves the password. Required [flags] -name/-n",
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
		file.Save(Name, data, key, Store)
	},
}

func init() {
	rootCmd.AddCommand(SaveCmd)
	SaveCmd.Flags().StringVarP(&Name, "name", "n", "", "Name of password.(Name will not be encrypted) Use list cmd to get Names")
	if err := SaveCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
}
