package leasing

import (
	"github.com/graphql-go/graphql"
	"strconv"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"leasing_contract": &graphql.Field{
			Type: LeasingType,
			Args: graphql.FieldConfigArgument{
				"leasing_id": &graphql.ArgumentConfig{
					Description: "Leasing ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["leasing_id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetLeasingContractByID(id)
			},
		},
	},
})