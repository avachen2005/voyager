package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	// "github.com/kr/pretty"
)

func String() string {

	if err := filepath.Walk(".", func(file string, info os.FileInfo, err error) error {

		var ret string = ""
		if strings.HasSuffix(file, ".graphql") {

			if b, err := ioutil.ReadFile(file); err != nil {
				panic(err)
			} else {
				ret = ret + strings.Trim(string(b))
			}
		}

		fmt.Println(ret)
		return nil
	}); err != nil {
		panic(err)
	}

	return ""
}

func main() {

	String()
}
