package graphql

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/packr"
)

func GenerateSchema() string {

	var ret string
	box := packr.NewBox(".")

	box.Walk(func(path string, f packr.File) error {

		if strings.HasSuffix(path, ".graphql") {
			ret = ret + box.String(path)
		}

		return nil
	})

	return ret
}
