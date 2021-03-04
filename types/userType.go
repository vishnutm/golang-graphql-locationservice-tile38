package types

import (
	"github.com/graphql-go/graphql"
)

// UserLocation : The structure for creating user's response [mutation:createUser]
type UserLocation struct {
	ID       string  `json:"id"`
	Lat      float64 `json:"latitude"`
	Long     float64 `json:"longitude"`
	Addressline1  string  `json:"address"`
	Street string `json:"street"`
	City    string `json:"city"`
	State    string   `json:"state"`
	
}

// GetLocation : The structure for querying user location response [query:user]
type GetLocation struct {
	LocationType string  `json:"type"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

// NearbyLocation : The structure for querying nearby users response [query:nearbyUser]
type NearbyLocation struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// UserType is the GraphQL schema for the response of creating user.
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserLocation",
	Fields: graphql.Fields{

		"id": &graphql.Field{
			Type: graphql.String,
		},
		"addressline1": &graphql.Field{
			Type: graphql.String,
		},
		"street": &graphql.Field{
			Type: graphql.String,
		},
		"city": &graphql.Field{
			Type: graphql.String,
		},
		"state": &graphql.Field{
			Type: graphql.String,
		},
		"latitude": &graphql.Field{
			Type: graphql.Float,
		},

		"longitude": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

// GetResponse is the GraphQL schema for the response of tile38 GET query
var GetResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetLocation",
	Fields: graphql.Fields{

		"type": &graphql.Field{
			Type: graphql.String,
		},
		"latitude": &graphql.Field{
			Type: graphql.Float,
		},
		"longitude": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

// NearbyType is the Graphql schema the response of tile38 NEARBY query
var NearbyType = graphql.NewObject(graphql.ObjectConfig{

	Name: "NearbyLocation",
	Fields: graphql.Fields{

		"id": &graphql.Field{
			Type: graphql.String,
		},
		"latitude": &graphql.Field{
			Type: graphql.Float,
		},
		"longitude": &graphql.Field{
			Type: graphql.Float,
		},
	},
})
