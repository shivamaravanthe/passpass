package file

import (
	"fmt"
	"io"
	"os"
	"strings"

	cryptopwd "github.com/shivamaravanthe/passpass/crypto"
)

func View(name string, key string, loc string) {
	f, err := os.Open(fmt.Sprintf("%s/%s.txt", loc, name))
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			fmt.Println("Requested password not saved use save cmd to save the password")
			return
		}
		fmt.Println("While Opening file ", err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("While Closing the file", err)
		}
	}()

	byteData, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("While Reading the file ", err)
		return
	}

	if len(byteData) == 0 {
		fmt.Printf("The Data is missing use update cmd to add data\n")
	} else {
		data, err := cryptopwd.Decrypt(byteData, key)
		if err == nil {
			fmt.Println(string(data))
		} else {
			fmt.Println("Wrong Key")
		}
	}
}
