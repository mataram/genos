package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/mataram/genos/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
