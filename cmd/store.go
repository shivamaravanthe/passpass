/*
Copyright © 2024 Shivanand Shenoy <vshivanand2@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var long string = `All the password saved in folder named store, 
store directory is required for working this application
create new store directory and set the location to store 
in ENV STORE_LOCATION
		
Create a Persistent Global Variable in Linux
https://www.freecodecamp.org/news/how-to-set-an-environment-variable-in-linux

Create a Persistent Global Variable in Windows
https://superuser.com/questions/949560/how-do-i-set-system-environment-variables-in-windows-10
`

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Gives the location of store file",
	Long:  long,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("File Location of store directory in STORE_LOCATION env variable is set as : ", Store)
		if Store == "" {
			fmt.Println(long)
		}
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
