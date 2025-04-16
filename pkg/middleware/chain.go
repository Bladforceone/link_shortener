package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := range middlewares {
			next = middlewares[i](next)
		}
		return next
	}
}
