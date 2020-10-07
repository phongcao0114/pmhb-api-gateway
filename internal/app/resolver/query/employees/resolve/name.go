package resolve

import (
	"fmt"
	"pmhb-api-gateway/models"
	"pmhb-api-gateway/models/parent"

	"github.com/graphql-go/graphql"
)

func GetName(p graphql.ResolveParams) (interface{}, error) {
	id := parent.ID(p)
	fmt.Println("-----------\nCall to Name service. id=", id)
	name := models.Name{
		FirstName: "John",
		LastName:  "Brown",
	}
	return name, nil
}
