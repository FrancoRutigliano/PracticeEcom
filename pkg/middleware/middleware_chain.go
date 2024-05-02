package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middleware ...Middleware) Middleware {
	// La función MiddlewareChain toma una lista variable de middlewares
	// y devuelve un solo middleware que representa la cadena completa.

	return func(next http.Handler) http.HandlerFunc {
		// Retorna un nuevo middleware que toma un http.Handler y devuelve un http.HandlerFunc.

		for i := len(middleware) - 1; i >= 0; i-- {
			// Itera sobre la lista de middlewares en orden inverso.

			next = middleware[i](next)
			// Aplica el middleware actual a la función next (que representa el siguiente middleware o el manejador final en la cadena).
		}

		return next.ServeHTTP
		// Retorna la función ServeHTTP del último middleware en la cadena, que será la función de cierre final
		// que se utilizará para procesar las solicitudes HTTP.
	}
}
