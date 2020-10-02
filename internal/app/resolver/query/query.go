package query

import (
	"github.com/graphql-go/graphql"
)

func Query() *graphql.Object {
	fields := graphql.Fields{
		"placeholder": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Boolean),
			Description: "Placeholder query",

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return true, nil
			},
		},

		"book_by_id": BookByID,
		"books":      Books,
		"test":       Test,
	}

	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: fields,
		},
	)
}
