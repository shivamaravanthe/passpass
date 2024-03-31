# passpass

passpass is a Golang Cli application to store passwords or any sensitive text information.

## Overview
* The Golang cli features are built using [Cobra](https://github.com/spf13/cobra) and [PromptUI](github.com/manifoldco/promptui) libraries.
* Data is encrypted using a user-entered key by using [AES](https://medium.com/@mertkimyonsen/encrypt-a-file-using-go-f1fe3bc7c635)  encryption and stored in a file in a folder named store.
* To view the Data same key must be entered.
* The store file can be in any location, its location must be stored in a global env variable STORE_LOCATION.
* The bin folder in this repository contains Linux binary and Windows exe files.


## Install Linux
Copy the passpass binary to /usr/local/bin

## Install Windows
Copy the passpass.exe in the folder mentioned in PATH env or add the location of passpass.exe to PATH
``` shell
> echo %PATH%
```
## Usage
``` shell
$ passpass
A simple password Vault, each password needs a master key to access

Usage:
  passpass [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List all the saved passwords
  save        Saves the password
  store       Gives the location of store file
  update      Updates the password
  view        Displays the Password

Flags:
  -h, --help   help for passpass

Use "passpass [command] --help" for more information about a command.

```
