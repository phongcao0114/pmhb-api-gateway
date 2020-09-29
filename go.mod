module pmhb-api-gateway

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/denisenkom/go-mssqldb v0.0.0-20200910202707-1e08a3fab204 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/graphql-go/graphql v0.7.9
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/phongcao0114/book-api-gateway v0.0.0-20200922033116-1e5acd5181b3 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	pmhb-book-service v0.0.0-20200929093511-dcb3a58670bf
)

replace pmhb-book-service v0.0.0-20200929093511-dcb3a58670bf => github.com/phongcao0114/pmhb-book-service v0.0.0-20200929093511-dcb3a58670bf
