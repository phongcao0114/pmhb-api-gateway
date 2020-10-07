package employees

import (
	"fmt"
	"pmhb-api-gateway/internal/app/datatype"
	"pmhb-api-gateway/models"

	"github.com/graphql-go/graphql"
)

var (
	Employees = &graphql.Field{
		Type: datatype.Employee,

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Println("------------------\ninside Employees resolver")
			employee := models.Employee{
				EmployeeID: "id001",
				Position:   "manager",
			}
			return employee, nil
		},
	}
)
