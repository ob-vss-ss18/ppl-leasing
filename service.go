package main

import (
	"fmt"
	"log"
	"net/http"
	//"encoding/json"
	"github.com/graphql-go/graphql"
	"os"
	"github.com/graphql-go/handler"
)

// The recommended way to use external packages is by downloading and installing them via "go get".
// the line "github.com/graphql-go/graphql"  is NOT an URL - its a folder structure.
// examples:
// https://github.com/topliceanu/graphql-go-example/blob/master/main.go
// https://gist.github.com/sogko/7debd336118e5e7c7f65


type Kunde struct {
	Kunden_id int `json:"kunden_id"`
	Name string `json:"name"`
	Vorname string `json:"vorname"`
}

var kunden_slice [] Kunde
var KundeType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Kunde",
		Fields: graphql.Fields{
			"kunden_id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				/*Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if kunde, ok := p.Source.(*Kunde); ok == true {
						return kunde.kunden_id, nil
					}
					return nil, nil
				},*/
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				/*Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if kunde, ok := p.Source.(*Kunde); ok == true {
						return kunde.name, nil
					}
					return nil, nil
				},*/
			},
			"vorname": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				/*Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if kunde, ok := p.Source.(*Kunde); ok == true {
						return kunde.vorname, nil
					}
					return nil, nil
				},*/
			},
		},
	},
)


type leasing_contract struct{
	Kunden_id int `json:"kunden_id"`
	Products []int `json:"products"`
	Leasing_id int  `json:"leasing_id"`
	Datum string `json:"datum"`
	Rabatt int `json:"rabatt"`
	Service_flat bool `json:"service_flat"`
	Testwert int `json:"testwert"`
	Versicherung bool `json:"versicherung"`

}

var leasing_contract_example leasing_contract = leasing_contract{
	1,
	[]int{1,2,3},
	0,
	"03.10.1992",
	4,
	true,
	10,
	false,
}

var leasing_type = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "leasingContract",
		Fields: graphql.Fields{
			"kunde": &graphql.Field{
				Type: KundeType,
			},
			"products": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
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
			"kunden" : &graphql.Field{
				Type: KundeType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if(len(kunden_slice)>0) {
						return &kunden_slice[0],nil
					};
					return nil,nil
				},
			},
			"all_kunden" : &graphql.Field{
				Type: KundeType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return &Kunde{
						7,
						"nachname",
						"vorname",
					}, nil;
				},
			},
			"liste" : &graphql.Field{
				Type: graphql.NewList(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					primes := []int{}
					primes = append(primes, 1)
					primes = append(primes, 234)
					return primes,nil
				},
			},
		},
	})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createKunde": &graphql.Field{
			Type: KundeType,
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
				user := Kunde{
					Kunden_id:len(kunden_slice),
					Name:nname,
					Vorname:vname,
				}
				kunden_slice = append(kunden_slice,user)
				return &user,nil
			},
		},
	},
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

//http://localhost:8080/graphql?query={user(id:"1"){name}}
//http://localhost:8080/graphql?query=mutation{createKunde(name: "hans", vorname: "dieter" ){kunden_id}}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, Dev Heroku!") // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName)                     // set router

	//_ = importJSONDataFromFile("data.json", &data)

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
		Query: rootQuery,
		Mutation: rootMutation,
	})



	handler := handler.New(&handler.Config{
		Schema:   &schema,
		GraphiQL: true,
	})


	/*http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		//check here for mutations
		var query = r.URL.Query().Get("query")
		var result *graphql.Result
		result = executeQuery(query, schema)
		fmt.Print("result received")
		json.NewEncoder(w).Encode(result)
	})*/

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.Handle("/graphql" , handler)



	err := http.ListenAndServe(":"+port, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
