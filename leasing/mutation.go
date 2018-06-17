package leasing

import (
	"github.com/graphql-go/graphql"
	"strconv"
	"strings"
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
					Type: 		 graphql.NewList(graphql.String),
				},
				"rabatt": &graphql.ArgumentConfig{
					Description: "Rabatt",
					Type: 		 graphql.NewList(graphql.Int),
				},
				"service_flat": &graphql.ArgumentConfig{
					Description: "Bool for the service flat",
					Type: 		 graphql.NewList(graphql.Boolean),
				},
				"testwert": &graphql.ArgumentConfig{
					Description: "Testwert",
					Type: 		 graphql.NewList(graphql.Boolean),
				},
				"versicherung": &graphql.ArgumentConfig{
					Description: "Bool for the insurance",
					Type: 		 graphql.NewList(graphql.Boolean),
				},
			},
			Resolve: func(p graphql.ResolveParams)(interface{}, error){
				k := p.Args["kunden_id"].(string)
				kunden_id, err := strconv.Atoi(k)
				if err != nil {
					return nil, err
				}
				prods := p.Args["products"].(string)
				//String in Liste umformen
				products := []int{};
				values := strings.Split(prods, ",")
				for product := range values {
					products = append(products, product)
				}

				datum := p.Args["datum"].(string)
				r := p.Args["rabatt"].(string)
				rabatt, errR := strconv.Atoi(r)
				if errR != nil {
					return nil, errR
				}
				s := p.Args["service_flat"].(string)
				service_flat, errS := strconv.ParseBool(s)
				if errS != nil {
					return nil, errS
				}
				t := p.Args["testwert"].(string)
				testwert, errT := strconv.ParseBool(t)
				if errT != nil {
					return nil, errT
				}
				v := p.Args["versicherung"].(string)
				versicherung, errV := strconv.ParseBool(v)
				if errV != nil {
					return nil, errV
				}
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
