package grifts

import (
	"github.com/cleisommais/oauth-service-v1/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
