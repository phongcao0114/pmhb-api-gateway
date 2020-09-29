package query

import (
	"encoding/json"
	"pmhb-api-gateway/internal/app/datatype"
	"pmhb-api-gateway/internal/pkg/khttp"

	"pmhb-book-service/models"

	"github.com/graphql-go/graphql"
	//"github.com/phongcao0114/pmhb-book-service/internal/app/models"
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
		id, _ := p.Args["id"].(string)
		url := "http://localhost:8080/kph/api/book/" + id
		header := map[string]string{
			"Content-Type": "application/json",
		}
		httpCaller := khttp.New(url, nil, header)
		resp, err := httpCaller.GET()
		if err != nil {
			return nil, err
		}
		var book models.Book
		json.Unmarshal(resp, &book)
		return book, nil
	},
}
