package datatype

import (
	employees "pmhb-api-gateway/internal/app/resolver/query/employees/resolve"

	"github.com/graphql-go/graphql"
)

var (
	Employee = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Employee",
			Fields: graphql.Fields{
				"employee_id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type:    Name,
					Resolve: employees.GetName,
				},
				"address": &graphql.Field{
					Type:    Address,
					Resolve: employees.Address,
				},
			},
		},
	)

	Name = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Name",
			Fields: graphql.Fields{
				"first_name": &graphql.Field{
					Type: graphql.String,
				},
				"last_name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	Address = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Address",
			Fields: graphql.Fields{
				"ward": &graphql.Field{
					Type: graphql.String,
				},
				"district": &graphql.Field{
					Type: graphql.String,
				},
				"province": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
)
