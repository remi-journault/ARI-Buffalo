package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/project/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
