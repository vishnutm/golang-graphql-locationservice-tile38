package mutations

import (
	//"encoding/json"
	"log"
	"net/http"
	c "pack/config"
	Error "pack/errors"
	"pack/types"
	//"strconv"
	//"time"

	"github.com/graphql-go/graphql"
	redis "github.com/go-redis/redis/v7"
)

// CreateUserMutation creates a new user in tile38 and returns the user details.
func CreateUserMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"addressline1":&graphql.ArgumentConfig{
				Type: graphql.String,
				DefaultValue: nil,
			},

			"street":&graphql.ArgumentConfig{
				Type: graphql.String,
				DefaultValue: nil,
			},
			"city":&graphql.ArgumentConfig{
				Type: graphql.String,
				DefaultValue: nil,
			},
			"state":&graphql.ArgumentConfig{
				Type: graphql.String,
				DefaultValue: nil,
			},

			"latitude": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},

			"longitude": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
		},

		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			var userlocation types.UserLocation
			userlocation.ID = params.Args["id"].(string)
			//userlocation.UserRole = params.Args["userrole"].(string)
			userlocation.Addressline1 = params.Args["address1"].(string)
			userlocation.Street = params.Args["street"].(string)
			userlocation.City = params.Args["city"].(string)
			userlocation.State = params.Args["state"].(string)
			userlocation.Lat = params.Args["latitude"].(float64)
			userlocation.Long = params.Args["longitude"].(float64)
			// tile38 command to create a new user.
			// The 'NX' keyword prevents creating new user with existing id.
			cmd := redis.NewStringCmd("SET",  userlocation.ID, "POINT", userlocation.Lat, userlocation.Long)
			c.Client.Process(cmd)
			_, err := cmd.Result()
			if err != nil {
				//log.Println(err)
				e:= Error.Wrap(err,"The User is aready existed",http.StatusConflict)
				log.Println(e)
				return e,e.Err
			}
			return userlocation, nil
		},
	}
}

// UpdateLocationMutation updates the location of an existing user.
func UpdateLocationMutation() *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"addressline1": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},

			"latitude": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},

			"longitude": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
		},

		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			ID := params.Args["id"].(string)
			UserRole := params.Args["userrole"].(string)
			Lat := params.Args["latitude"].(float64)
			Long := params.Args["longitude"].(float64)

			// tile38 command to update an existing user.
			// The 'XX' keyword updates location only if the id already exist else returns error instead of creating new user.
			cmd := redis.NewStringCmd("SET", UserRole, ID, "XX", "POINT", Lat, Long)
			c.Client.Process(cmd)
			v, err := cmd.Result()
			if err != nil {
				//log.Println(err)
				e:= Error.Wrap(err,"User not found for updating location of user",http.StatusNotFound)
				log.Println(e)
				return e,e.Err
			}
			return v, nil
		},
	}
}



