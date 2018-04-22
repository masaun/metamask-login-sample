package auth

import (
	"github.com/m0t0k1ch1/metamask-login-sample/application/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

func SetUp(cntl *server.Controller) {
	cntl.POST("/challenge", ChallengeHandler)
	cntl.POST("/authorize", AuthorizeHandler)
}

func ChallengeHandler(c *server.Context) error {
	addressHex := c.FormValue("address")

	app := c.NewAuthApplication()

	ctx := c.Request().Context()
	in := auth.NewChallengeInput(addressHex)

	out, err := app.Challenge(ctx, in)
	if err != nil {
		return c.JSONError(err)
	}

	return c.JSONSuccess(out)
}

func AuthorizeHandler(c *server.Context) error {
	addressHex := c.FormValue("address")
	sigHex := c.FormValue("signature")

	app := c.NewAuthApplication()

	ctx := c.Request().Context()
	in := auth.NewAuthorizeInput(addressHex, sigHex)

	out, err := app.Authorize(ctx, in)
	if err != nil {
		return c.JSONError(err)
	}

	return c.JSONSuccess(out)
}
