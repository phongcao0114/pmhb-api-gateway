package mutation

import (
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/app/validation/header"

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
		//Validate token
		if err := header.ValidateJWT(p.Context); err != nil {
			return nil, err
		}

		//Parse arguments
		id, _ := p.Args["id"].(string)

		//Make HTTP call
		url := config.Config.GraphQLServicePath.BookService + "/kph/api/book/" + id
		httpCaller := utils.MakeHTTPCaller(url, nil)
		resp, err := httpCaller.DELETE()
		if err != nil {
			return nil, err
		}

		//Handle response from service
		var status bool
		return utils.HandleResp(resp, &status)
	},
}
