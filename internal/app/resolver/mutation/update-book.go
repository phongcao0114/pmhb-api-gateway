package mutation

import (
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/pkg/khttp"
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
		id, _ := p.Args["id"].(string)
		name, _ := p.Args["name"].(string)
		author, _ := p.Args["author"].(string)
		bookReq := models.UpdateBookReq{
			Name:   name,
			Author: author,
		}

		url := config.Config.GraphQLServicePath.BookService + "/kph/api/book/" + id
		header := map[string]string{
			"Content-Type": "application/json",
		}
		httpCaller := khttp.New(url, bookReq, header)
		_, err := httpCaller.PUT()
		if err != nil {
			return nil, err
		}
		return true, nil
	},
}
