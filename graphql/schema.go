package graphql

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GenerateSchema() string {

	var ret string = ""

	if err := filepath.Walk(".", func(file string, info os.FileInfo, err error) error {
		if strings.HasSuffix(file, ".graphql") {

			if b, err := ioutil.ReadFile(file); err != nil {
				panic(err)
			} else {
				ret = ret + strings.TrimSpace(string(b))
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}

	return ret
}
