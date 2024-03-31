package file

import (
	"fmt"
	"os"
	"strings"
)

func List(loc string) {
	f, err := os.ReadDir(loc)
	if err != nil {
		fmt.Println("While Opening store folder", err)
		return
	}

	for _, file := range f {
		fmt.Println("->", strings.Split(file.Name(), ".txt")[0])
	}
}
