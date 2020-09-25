package api

import (
	"net/http"
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/handlers"
)

type (
	middleware = func(http.Handler) http.Handler
	route      struct {
		desc        string
		path        string
		method      string
		handler     http.HandlerFunc
		middlewares []middleware
	}
)

// CreateGraphQLHandler function
func CreateGraphQLHandler(conf *config.Configs) *handlers.GraphQLHandler {
	return handlers.NewGraphQLHandler(conf)
}
