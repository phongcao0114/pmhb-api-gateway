package query

import (
	"pmhb-api-gateway/internal/app/datatype"

	"github.com/graphql-go/graphql"
)

var Books = &graphql.Field{
	Type:        datatype.ListBook,
	Description: "Get book list",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}

//token := p.Context.Value("token").(string)
//err := header.ValidateJWT(token)
//if err != nil {
//	return nil, err
//}
//url := "http://localhost:8080/api/v1/books"
//resp, err := helper.MakeRequest(http.MethodGet, url, nil)
//if err != nil {
//	return []model.Book{}, err
//}
//defer resp.Body.Close()
//
//var books []model.Book
//json.NewDecoder(resp.Body).Decode(&books)
//return books, nil
