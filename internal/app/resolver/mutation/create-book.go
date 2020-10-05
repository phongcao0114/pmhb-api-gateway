package mutation

import (
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/pkg/khttp"

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
		name, _ := p.Args["name"].(string)
		author := p.Args["author"].(string)
		bookReq := models.InsertBookReq{
			Name:   name,
			Author: author,
		}

		url := config.Config.GraphQLServicePath.BookService + "/kph/api/book"
		header := map[string]string{
			"Content-Type": "application/json",
		}
		httpCaller := khttp.New(url, bookReq, header)
		resp, err := httpCaller.POST()
		if err != nil {
			return nil, err
		}
		var id string
		return utils.HandleResp(resp, &id)
	},
}
