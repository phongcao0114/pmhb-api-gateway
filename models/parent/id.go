package parent

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

type IDer interface {
	ID() string
}

func ID(p graphql.ResolveParams) string {
	fmt.Println("--------\ninside parentid")
	if s, ok := p.Source.(IDer); !ok {
		fmt.Println("--------\nparent: ID: source is not IDer")
		return ""

	} else {
		return s.ID()
	}
}
