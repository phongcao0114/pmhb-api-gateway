package query

import (
	"encoding/json"
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/datatype"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/pkg/khttp"
	"pmhb-book-service/models"

	"github.com/graphql-go/graphql"
)

var Books = &graphql.Field{
	Type:        datatype.ListBook,
	Description: "Get book list",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		url := config.Config.GraphQLServicePath.BookService + "/kph/api/book"
		header := map[string]string{
			"Content-Type": "application/json",
		}
		httpCaller := khttp.New(url, nil, header)
		resp, err := httpCaller.GET()
		if err != nil {
			return nil, err
		}
		err = utils.HandleRespError(resp)
		if err != nil {
			return nil, err
		}
		var books []models.Book

		err = json.Unmarshal(resp, &books)
		if err != nil {
			return nil, err
		}
		return books, nil
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
