package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"strconv"
	"os"
)

// The recommended way to use external packages is by downloading and installing them via "go get".
// the line "github.com/graphql-go/graphql"  is NOT an URL - its a folder structure.
// examples:
// https://github.com/topliceanu/graphql-go-example/blob/master/main.go
// https://gist.github.com/sogko/7debd336118e5e7c7f65


type Kunde struct {
	kunden_id int
	name string
	vorname string
}

var kunden_slice []Kunde
var kunde_type = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Kunde",
		Fields: graphql.Fields{
			"kunden_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if kunde, ok := p.Source.(*Kunde); ok == true {
						return kunde.kunden_id, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if kunde, ok := p.Source.(*Kunde); ok == true {
						return kunde.name, nil
					}
					return nil, nil
				},
			},
			"vorname": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if kunde, ok := p.Source.(*Kunde); ok == true {
						return kunde.vorname, nil
					}
					return nil, nil
				},
			},
		},
	},
)

type leased_products struct {
	product_ids []int
	leasing_id int
}

var leased_products_type = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "leasedProducts",
		Fields: graphql.Fields{
			"product_ids": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"leasing_id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

type leasing_contract struct{
	kunde Kunde
	products []leased_products
	datum string
	rabatt int
	service_flat bool
	peak_flat bool
	testwert int
	versicherung bool

}

var leasing_type = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "leasingContract",
		Fields: graphql.Fields{
			"kunde": &graphql.Field{
				Type: kunde_type,
			},
			"products": &graphql.Field{
				Type: graphql.NewList(leased_products_type),
			},
			"datum": &graphql.Field{
				Type: graphql.String,
			},
			"rabatt": &graphql.Field{
				Type: graphql.Int,
			},
			"service_flat": &graphql.Field{
				Type: graphql.Boolean,
			},
			"peak_flat": &graphql.Field{
				Type: graphql.Boolean,
			},
			"testwert": &graphql.Field{
				Type: graphql.Int,
			},
			"versicherung": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

/*
   Create Query object type with fields "user" has type [userType] by using GraphQLObjectTypeConfig:
       - Name: name of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig to define:
       - Type: type of field
       - Args: arguments to query with current field
       - Resolve: function to query data using params from [Args] and return value with current type
*/

 var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"test" : &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"value": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams)(interface{}, error){
					var test = ""
					test = "test successfull - your input was: " + p.Args["value"].(string)
					return test,nil
				},
			},
			"test2" : &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"value": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"rand": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams)(interface{}, error){
					var test = ""
					if(len(p.Args) == 0){
						return "no patameter found", nil
					}
					if(p.Args["value"] == nil){
						test += "value not found. "
					}else{
						test += "value was: " + strconv.Itoa(p.Args["value"].(int)) + ". "
					}
					if(p.Args["rand"] == nil){
						test += "rand not found. "
					}else{
						test += "rand was: " + strconv.Itoa(p.Args["rand"].(int)) + ". "
					}
					return test,nil
				},
			},
			"kunden" : &graphql.Field{
				Type: kunde_type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if(len(kunden_slice)>0) {
						return &kunden_slice[0],nil
					};
					return nil,nil
				},
			},
		},
	})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createKunde": &graphql.Field{
			Type: kunde_type,
			Args: graphql.FieldConfigArgument{
				"vorname": &graphql.ArgumentConfig{
					Description: "user forename",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Description: "user surname",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams)(interface{}, error){
				nname := p.Args["name"].(string)
				vname := p.Args["vorname"].(string)
				user := &Kunde{
					vorname:vname,
					name:nname,
					kunden_id:len(kunden_slice),
				}
				kunden_slice = append(kunden_slice,*user)
				return user,nil
			},
		},
	},
})


var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
		Mutation: rootMutation,
	})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

//http://localhost:8080/graphql?query={user(id:%221%22){name}}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, Dev Heroku!") // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName)                     // set router

	//_ = importJSONDataFromFile("data.json", &data)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		//check here for mutations
		var query = r.URL.Query().Get("query")
		var result *graphql.Result
		result = executeQuery(query, schema)
		fmt.Print("result received")
		json.NewEncoder(w).Encode(result)
	})

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil) // set listen port
	//err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
