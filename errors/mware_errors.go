package errors

import (
	"github.com/praesarium/go-engine/engine"
	"net/http"
	"strings"
)

func MiddlewareErrors(renderer engine.Renderer) engine.Middleware {
	var httpStatusCode string
	var httpMessage    string

	return func(c *engine.Context) {
		c.NextMiddleware()

		httpStatusCode = c.Writer.Code()
		if httpStatusCode > http.StatusOK {
			httpMessage = strings.ToLower(http.StatusText(httpStatusCode))

			// render error message
			c.Render(httpStatusCode, renderer,
				engine.H{
					"message": httpMessage,
				},
			)
		}
	}
}
