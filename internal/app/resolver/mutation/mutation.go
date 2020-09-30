package mutation

import (
	"github.com/graphql-go/graphql"
)

func Mutation() *graphql.Object {
	fields := graphql.Fields{
		"placeholder": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Boolean),
			Description: "Placeholder mutation",

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return true, nil
			},
		},

		"create_book": CreateBook,
		"delete_book": DeleteBook,
		"update_book": UpdateBook,
	}

	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: fields,
		},
	)
}
