package mutation

import (
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/pkg/khttp"

	"github.com/graphql-go/graphql"
)

var DeleteBook = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "DeleteBook",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(string)
		url := config.Config.GraphQLServicePath.BookService + "/kph/api/book/" + id
		header := map[string]string{
			"Content-Type": "application/json",
		}
		httpCaller := khttp.New(url, nil, header)
		resp, err := httpCaller.DELETE()
		if err != nil {
			return nil, err
		}
		var status bool
		return utils.HandleResp(resp, &status)
	},
}
