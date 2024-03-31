package file

import (
	"fmt"
	"io"
	"os"
	"strings"

	cryptopwd "github.com/shivamaravanthe/passpass/crypto"
)

func Update(name string, newName string, newData string, key string, loc string) {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s.txt", loc, name), os.O_RDWR, os.ModeAppend)
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

	data, err := cryptopwd.Decrypt(byteData, key)
	if err == nil {
		fmt.Println("Previous Data", string(data))
	} else {
		fmt.Println("Wrong Key")
		return
	}

	NewByteData, err := cryptopwd.Encrypt([]byte(newData), key)
	if err == nil {
		if _, err := f.Seek(0, 0); err != nil {
			fmt.Println("While Seek file", err)
			return
		}
		if err := f.Truncate(0); err != nil {
			fmt.Println("While Truncate file", err)
			return
		}
		if _, err = f.Write(NewByteData); err != nil {
			fmt.Println("While Writing data", err)
			return
		}
	}

	if newName == "" {
		return
	}
	files, err := os.ReadDir(loc)
	if err != nil {
		fmt.Println("While Opening store folder", err)
		return
	}

	for _, file := range files {
		if file.Name() == fmt.Sprintf("%s.txt", newName) {
			fmt.Printf("Password Name %s already exist Choose different name or update the same\n", newName)
			return
		}
	}
	if err := os.Rename(fmt.Sprintf("%s/%s.txt", loc, name), fmt.Sprintf("%s/%s.txt", loc, newName)); err != nil {
		fmt.Println("While Renaming file", err)
		return
	}
}
