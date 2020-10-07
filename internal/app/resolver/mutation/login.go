package mutation

import (
	"errors"
	"pmhb-api-gateway/internal/app/utils"
	"pmhb-api-gateway/internal/app/validation/header"

	"golang.org/x/crypto/bcrypt"

	"github.com/common-go/jwt"
	"github.com/graphql-go/graphql"
)

var Login = &graphql.Field{
	Type:        graphql.String,
	Description: "Login",
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//Parse arguments
		username := p.Args["username"].(string)
		password := p.Args["password"].(string)

		//Handle login
		hashedPasswordFromDB, _ := hashPassword("P@ssw0rd") //TODO: get from DB
		flag := checkPasswordHash(password, hashedPasswordFromDB)
		if !flag {
			return nil, errors.New("wrong password")
		}

		tokenService := jwt.DefaultTokenService{}
		exp := int64(15 * 60 * 1000)
		payload := utils.JWTPayload{
			Username: username,
		}

		//Generate token
		token, err := tokenService.GenerateToken(payload, header.SecretKey, exp)
		if err != nil {
			return nil, err
		}
		return token, nil
	},
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
