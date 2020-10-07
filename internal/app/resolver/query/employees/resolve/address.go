package resolve

import (
	"fmt"
	"pmhb-api-gateway/models"
	"pmhb-api-gateway/models/parent"

	"github.com/graphql-go/graphql"
)

func Address(p graphql.ResolveParams) (interface{}, error) {
	id := parent.ID(p)
	fmt.Println("-----------\nCall to Address service. id=", id)
	address := models.Address{
		Ward:     "ward1",
		District: "district1",
		Province: "province1",
	}
	return address, nil
}
