package leasing

import (
	"github.com/graphql-go/graphql"
	"strconv"
	"time"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createLeasing": &graphql.Field{
			Type: LeasingType,
			Args: graphql.FieldConfigArgument{
				"kunden_id": &graphql.ArgumentConfig{
					Description: "Kunden ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"products": &graphql.ArgumentConfig{
					Description: "List with products",
					Type: 		 graphql.NewList(graphql.Int),
				},
				"datum": &graphql.ArgumentConfig{
					Description: "Datum",
					Type: 		 graphql.NewNonNull(graphql.DateTime),
				},
				"rabatt": &graphql.ArgumentConfig{
					Description: "Rabatt",
					Type: 		 graphql.NewNonNull(graphql.Int),
				},
				"service_flat": &graphql.ArgumentConfig{
					Description: "Bool for the service flat",
					Type: 		 graphql.NewNonNull(graphql.Boolean),
				},
				"testwert": &graphql.ArgumentConfig{
					Description: "Testwert",
					Type: 		 graphql.NewNonNull(graphql.Boolean),
				},
				"versicherung": &graphql.ArgumentConfig{
					Description: "Bool for the insurance",
					Type: 		 graphql.NewNonNull(graphql.Boolean),
				},
			},
			Resolve: func(p graphql.ResolveParams)(interface{}, error){
				k := p.Args["kunden_id"].(string)
				kunden_id, err := strconv.Atoi(k)
				if err != nil {
					return nil, err
				}
				prods, _ := p.Args["products"].([]interface{})
				products := []int{}
				for _, product := range prods {
					products = append(products, product.(int))
				}
				datumString := p.Args["datum"].(string)
				layout := "2006-01-02T15:04:05.000Z"
				datum, err := time.Parse(layout, datumString)
				if err != nil {
					return nil, err
				}
				rabatt := p.Args["rabatt"].(int)
				service_flat := p.Args["service_flat"].(bool)
				testwert := p.Args["testwert"].(bool)
				versicherung := p.Args["versicherung"].(bool)
				leasing_contract := Leasing_contract{
					Kunden_id: kunden_id,
					Products: products,
					Datum: datum,
					Rabatt: rabatt,
					Service_flat: service_flat,
					Testwert: testwert,
					Versicherung: versicherung,
				}
				return CreateLeasingContract(&leasing_contract)
			},
		},
	},
})
