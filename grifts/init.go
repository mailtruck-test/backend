package grifts

import (
  "github.com/gobuffalo/buffalo"
	"userreg/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
