package mutation

import (
	"errors"
	"pmhb-api-gateway/internal/app/utils"

	"golang.org/x/crypto/bcrypt"

	"github.com/common-go/jwt"
	"github.com/graphql-go/graphql"
)

var Login = &graphql.Field{
	Type:        graphql.String,
	Description: "Login",
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		username := p.Args["username"].(string)
		password := p.Args["password"].(string)

		hashedPasswordFromDB, _ := hashPassword("P@ssw0rd") //TODO: get from DB
		flag := checkPasswordHash(password, hashedPasswordFromDB)
		if !flag {
			return nil, errors.New("wrong password")
		}
		//Login is assumed always success
		tokenService := jwt.DefaultTokenService{}
		exp := int64(15 * 60 * 1000)
		payload := utils.JWTPayload{
			Username: username,
		}
		token, err := tokenService.GenerateToken(payload, utils.SecretKey, exp)
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
