package grifts

import (
	"github.com/avachen2005/voyager/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
