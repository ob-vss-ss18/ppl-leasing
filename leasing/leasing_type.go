package leasing

import "github.com/graphql-go/graphql"

var LeasingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "leasingContract",
	Fields: graphql.Fields{
		"leasing_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"kunden_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"products": &graphql.Field{
			Type: graphql.NewList(graphql.Int),
		},
		"datum": &graphql.Field{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
		"rabatt": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"service_flat": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"testwert": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"versicherung": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
},
)
