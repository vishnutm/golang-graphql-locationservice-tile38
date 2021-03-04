package queries

import (
	"net/http"
	"pack/types"

	// "encoding/json"
	"log"
	c "pack/config"
	Error "pack/errors"

	redis "github.com/go-redis/redis/v7"
	"github.com/graphql-go/graphql"
	"github.com/tidwall/gjson"
)

// GetUserQuery returns the latitude and longitude of the user with given 'id' and 'userRole'.
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: types.GetResponse,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"userRole": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"].(string)
			userRole := params.Args["userRole"].(string)
			cmd := redis.NewStringCmd("GET", userRole, id)
			// c.Client.Process(redis.NewStringCmd("OUTPUT", "json")) //redis: parse string
			c.Client.Process(cmd)
			v, err := cmd.Result()
			if err != nil {
				//log.Println(err)
				e := Error.Wrap(err, "Your requested User location is not found", http.StatusNotFound)
				log.Println(e)
				return e, e.Err
			}
			var response types.GetLocation
			response.LocationType = gjson.Get(v, "type").String()
			response.Latitude = gjson.Get(v, "coordinates.1").Float()
			response.Longitude = gjson.Get(v, "coordinates.0").Float()

			return response, nil
		},
	}
}

// GetNearbyUser returns a list of users and their location within a given radius of a given location
func GetNearbyUser() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.NearbyType),
		Args: graphql.FieldConfigArgument{

			"latitude": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},

			"longitude": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},

			"radius": &graphql.ArgumentConfig{
				Type:         graphql.Float,
				DefaultValue: 2000.00,
			},
		},

		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			var response []types.NearbyLocation

			var Lat = params.Args["latitude"].(float64)
			var Long = params.Args["longitude"].(float64)
			var radius = params.Args["radius"].(float64)
			// if params.Args["radius"] == nil  {
			// 	radius = params.Args["radius"].(float64)
			// }
			cmd := redis.NewStringCmd("NEARBY", "driver", "POINT", Lat, Long, radius)
			c.Client.Process(redis.NewStringCmd("OUTPUT", "json"))
			c.Client.Process(cmd)
			v, err := cmd.Result()
			if err != nil {
				//log.Println(err)
				e := Error.Wrap(err, "Your requested user is not found nearby", http.StatusNotFound)
				return e, e.Err
			}
			// log.Println(v)
			var res types.NearbyLocation
			result := gjson.Get(v, "objects")
			result.ForEach(func(key, value gjson.Result) bool {
				res.ID = gjson.Get(value.String(), "id").String()
				res.Latitude = gjson.Get(value.String(), "object.coordinates").Array()[1].Array()[0].Float()
				res.Longitude = gjson.Get(value.String(), "object.coordinates").Array()[0].Array()[0].Float()
				response = append(response, res)
				return true // keep iterating
			})
			return response, nil
		},
	}
}



// GetAllDrivers returns all the drivers in tile38 server
func GetAllDrivers() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		// Args: graphql.FieldConfigArgument{
		// 	"userRole":&graphql.ArgumentConfig{
		// 		Type:graphql.NewNonNull(graphql.String),
		// 	},
		// },

		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			var response []types.UserLocation

			// var fleet = params.Args["userRole"].(string)

			cmd := redis.NewStringCmd("SCAN", "driver")
			c.Client.Process(redis.NewStringCmd("OUTPUT", "json"))
			c.Client.Process(cmd)
			v, err := cmd.Result()
			if err != nil {
				log.Println(err)
				return nil, err
			}
			// log.Println(v)
			var res types.UserLocation
			result := gjson.Get(v, "objects")
			result.ForEach(func(key, value gjson.Result) bool {
				res.ID = gjson.Get(value.String(), "id").String()
				res.Lat = gjson.Get(value.String(), "object.coordinates").Array()[1].Array()[0].Float()
				res.Long = gjson.Get(value.String(), "object.coordinates").Array()[0].Array()[0].Float()
				response = append(response, res)
				return true // keep iterating
			})
			return response, nil
		},
	}
}
