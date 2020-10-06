package mutation

import (
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/app/validation/header"
	"pmhb-book-service/models"

	"github.com/graphql-go/graphql"
)

var CreateBook = &graphql.Field{
	Type:        graphql.String,
	Description: "CreateBook",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"author": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//Validate token
		if err := header.ValidateJWT(p.Context); err != nil {
			return nil, err
		}

		//Parse arguments
		name, _ := p.Args["name"].(string)
		author := p.Args["author"].(string)

		//Make HTTP call
		bookReq := models.InsertBookReq{
			Name:   name,
			Author: author,
		}
		url := config.Config.GraphQLServicePath.BookService + "/kph/api/book"
		httpCaller := utils.MakeHTTPCaller(url, bookReq)
		resp, err := httpCaller.POST()
		if err != nil {
			return nil, err
		}

		//Handle response from service
		var id string
		return utils.HandleResp(resp, &id)
	},
}
