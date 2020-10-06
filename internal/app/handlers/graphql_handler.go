package handlers

import (
	"context"
	"log"
	"net/http"
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/resolver/mutation"
	"pmhb-api-gateway/internal/app/resolver/query"
	"pmhb-api-gateway/internal/app/response"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/app/validation/header"
	"pmhb-api-gateway/internal/kerrors"
	"pmhb-api-gateway/internal/pkg/klog"

	"github.com/graphql-go/graphql"
)

const (
	// GraphQLHandlerPrefix prefix logger
	GraphQLHandlerPrefix = "GraphQL_handler"
)

// GraphQLHandler struct defines the variables for specifying interface.
type GraphQLHandler struct {
	conf       *config.Configs
	errHandler kerrors.KError
	logger     klog.Logger

	//srv services.TransactionsService
}

// NewGraphQLHandler func
func NewGraphQLHandler(conf *config.Configs) *GraphQLHandler {
	return &GraphQLHandler{
		conf:       conf,
		errHandler: kerrors.WithPrefix(GraphQLHandlerPrefix),
		logger:     klog.WithPrefix(GraphQLHandlerPrefix),
	}
}

// GraphqlHandler func
func (g *GraphQLHandler) GraphqlHandler(w http.ResponseWriter, r *http.Request) {
	var graphQLPostBody utils.GraphQLPostBody
	err := utils.DecodeToBody(&g.errHandler, &graphQLPostBody, r)
	if err != nil {
		response.WriteJSON(w)(response.HandleError(r, err))
		return
	}
	token, err := header.GetTokenFromHttpRequest(r)
	if err != nil {
		token = ""
	}

	result := graphql.Do(graphql.Params{
		Schema:         *Init(),
		RequestString:  graphQLPostBody.Query,
		VariableValues: graphQLPostBody.Variables,
		OperationName:  graphQLPostBody.OperationName,
		Context:        context.WithValue(context.Background(), header.TokenHeaderKey, token),
	})
	response.WriteJSON(w)(response.HandleSuccess(r, result))
	return
}

func Init() *graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    query.Query(),
			Mutation: mutation.Mutation(),
		},
	)

	if err != nil {
		log.Fatalf("schema: Init: %v", err)
		return nil
	}

	return &schema
}
