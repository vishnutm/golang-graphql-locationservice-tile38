package main

import (
	"fmt"
	"log"
	"net/http"

	// "encoding/json"
	c "pack/config"
	Error "pack/errors"
	"pack/mutations"
	"pack/queries"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/spf13/viper"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GraphQL APIs using graphql-go!")
}

func main() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")

	viper.SetConfigType("yml")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)

	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	router := mux.NewRouter().StrictSlash(true)
	log.Printf(" Server started on http://localhost:8080/\n")

	// GraphQl schema configuration
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(),
		}),
	}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	router.Handle("/v1", httpHandler)
	router.HandleFunc("/", homeLink)
	if _, err := c.Client.Ping().Result(); err != nil {

		e := Error.Wrap(err, ",Redis server down,", http.StatusBadGateway)
		//log.Println(e.Status)
		//	fmt.Println(e.Content ,e.Err,e.Status)
		log.Println(e)
	}
	log.Fatal(http.ListenAndServe(":"+configuration.Server.Port, router))

}
