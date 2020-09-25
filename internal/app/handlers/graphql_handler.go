package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/models"
	"pmhb-api-gateway/internal/app/resolver/query"
	"pmhb-api-gateway/internal/app/response"
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
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.WriteJSON(w)(response.HandleError(r, err))
			return
		}

		var graphQLPostBody models.GraphQLPostBody
		err = json.Unmarshal(body, &graphQLPostBody)
		if err != nil {
			response.WriteJSON(w)(response.HandleError(r, err))
			return
		}

		//token, _ := header.GetTokenFromHttpRequest(r)
		//if err != nil {
		//	//	response.WriteJSON(w)(response.HandleError(r, err))
		//	//	return
		//	token = ""
		//}
		result := graphql.Do(graphql.Params{
			Schema:         *Init(),
			RequestString:  graphQLPostBody.Query,
			VariableValues: graphQLPostBody.Variables,
			OperationName:  graphQLPostBody.OperationName,
			//Context:        context.WithValue(context.Background(), "token", token),
		})
		json.NewEncoder(w).Encode(result)

	default:
		fmt.Fprintf(w, "Sorry, only POST method are supported.")
	}
}

func Init() *graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: query.Query(),
			//Mutation: mutation.Mutation(),
		},
	)

	if err != nil {
		log.Fatalf("schema: Init: %v", err)
		return nil
	}

	return &schema
}
