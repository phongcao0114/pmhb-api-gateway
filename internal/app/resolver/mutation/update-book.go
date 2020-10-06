package mutation

import (
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/app/validation/header"
	"pmhb-book-service/models"

	"github.com/graphql-go/graphql"
)

var UpdateBook = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Update book",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"author": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//Validate token
		if err := header.ValidateJWT(p.Context); err != nil {
			return nil, err
		}

		//Parse arguments
		id, _ := p.Args["id"].(string)
		name, _ := p.Args["name"].(string)
		author, _ := p.Args["author"].(string)

		//Make HTTP call
		bookReq := models.UpdateBookReq{
			Name:   name,
			Author: author,
		}
		url := config.Config.GraphQLServicePath.BookService + "/kph/api/book/" + id
		httpCaller := utils.MakeHTTPCaller(url, bookReq)
		resp, err := httpCaller.PUT()
		if err != nil {
			return nil, err
		}

		//Handle response from service
		var status bool
		return utils.HandleResp(resp, &status)
	},
}
