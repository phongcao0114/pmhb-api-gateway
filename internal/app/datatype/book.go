package datatype

import "github.com/graphql-go/graphql"

var (
	ListBook = graphql.NewList(Book)
	Book     = graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "Book",
			Description: "Book",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"author": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
)
