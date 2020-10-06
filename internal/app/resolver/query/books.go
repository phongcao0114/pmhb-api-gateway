package query

import (
	"pmhb-api-gateway/internal/app/config"
	"pmhb-api-gateway/internal/app/datatype"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/pkg/khttp"
	"pmhb-book-service/models"

	"github.com/common-go/jwt"

	"github.com/graphql-go/graphql"
)

var Books = &graphql.Field{
	Type:        datatype.ListBook,
	Description: "Get book list",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//Validate token
		tokenService := jwt.DefaultTokenService{}
		tokenString := p.Context.Value("token").(string)
		_, _, _, err := tokenService.VerifyToken(tokenString, utils.SecretKey)
		if err != nil {
			return nil, err
		}
		//Make HTTP call
		url := config.Config.GraphQLServicePath.BookService + "/kph/api/book"
		header := map[string]string{
			"Content-Type": "application/json",
		}
		httpCaller := khttp.New(url, nil, header)
		resp, err := httpCaller.GET()
		if err != nil {
			return nil, err
		}

		//Handle response from service
		var books []models.Book
		return utils.HandleResp(resp, &books)
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
