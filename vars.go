package httprouter

import (
	"context"
	"net/http"
)

const (
	varsKey  = `httprouter_route_vars`
	routeKey = `httprouter_route`
)

// ContextVars returns the route variables for the current Context, if any.
func ContextVars(ctx context.Context) map[string]string {
	if rv := ctx.Value(varsKey); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

// Vars returns the route variables for the current request, if any.
func Vars(r *http.Request) map[string]string {
	if rv := contextGet(r, varsKey); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

func setVars(r *http.Request, val interface{}) *http.Request {
	return contextSet(r, varsKey, val)
}

func contextGet(r *http.Request, key interface{}) interface{} {
	return r.Context().Value(key)
}

func contextSet(r *http.Request, key, val interface{}) *http.Request {
	if val == nil {
		return r
	}

	return r.WithContext(context.WithValue(r.Context(), key, val))
}

func contextClear(r *http.Request) {
	return
}
