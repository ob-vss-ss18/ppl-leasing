package leasing

import (
	"github.com/graphql-go/handler"
	"github.com/graphql-go/graphql"
	"log"
)

var ApiHandler *handler.Handler

func Init() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		log.Fatal(err)
	}

	ApiHandler = handler.New(&handler.Config{
		Schema:   &schema,
		GraphiQL: true,
	})
}