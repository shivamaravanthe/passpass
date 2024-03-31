package file

import (
	"fmt"
	"io"
	"os"

	cryptopwd "github.com/shivamaravanthe/passpass/crypto"
)

func Save(name string, data string, key string, loc string) {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s.txt", loc, name), os.O_RDWR, os.ModeAppend)
	if err != nil {
		f, err = os.Create(fmt.Sprintf("%s/%s.txt", loc, name))
		if err != nil {
			fmt.Println("While Creating file", err)
			return
		}
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

	content := string(byteData)

	if len(content) != 0 {
		fmt.Printf("Password for %s Already present use update cmd to change value or view cmd to view the value\n", name)
		return
	} else {
		byteData, err := cryptopwd.Encrypt([]byte(data), key)
		if err == nil {
			if _, err = f.Write(byteData); err != nil {
				fmt.Println("While Writing data", err)
				return
			}
		}
	}
	fmt.Println("successfully saved")
}
