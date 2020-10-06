package query

import (
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/datatype"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/app/validation/header"

	"pmhb-book-service/models"

	"github.com/graphql-go/graphql"
)

var BookByID = &graphql.Field{
	Type:        datatype.Book,
	Description: "BookByID",
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
		resp, err := httpCaller.GET()
		if err != nil {
			return nil, err
		}

		//Handle response from service
		var book models.Book
		return utils.HandleResp(resp, &book)
	},
}
