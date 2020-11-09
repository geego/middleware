package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/geego/geen"
)

var (
	// URLFormatCtxKey is the context.Context key to store the URL format data
	// for a request.
	URLFormatCtxKey = &contextKey{"URLFormat"}
)

func URLFormat(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var format string
		path := r.URL.Path

		if strings.Index(path, ".") > 0 {
			base := strings.LastIndex(path, "/")
			idx := strings.Index(path[base:], ".")

			if idx > 0 {
				idx += base
				format = path[idx+1:]

				rctx := geen.RouteContext(r.Context())
				rctx.RoutePath = path[:idx]
			}
		}

		r = r.WithContext(context.WithValue(ctx, URLFormatCtxKey, format))

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
