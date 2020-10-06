package query

import (
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/datatype"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/app/validation/header"
	"pmhb-book-service/models"

	"github.com/graphql-go/graphql"
)

var Books = &graphql.Field{
	Type:        datatype.ListBook,
	Description: "Get book list",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//Validate token
		if err := header.ValidateJWT(p.Context); err != nil {
			return nil, err
		}

		//Make HTTP call
		url := config.Config.GraphQLServicePath.BookService + "/kph/api/book"
		httpCaller := utils.MakeHTTPCaller(url, nil)
		resp, err := httpCaller.GET()
		if err != nil {
			return nil, err
		}

		//Handle response from service
		var books []models.Book
		return utils.HandleResp(resp, &books)
	},
}
