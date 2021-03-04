package queries

import (
	//"pack/types"

	"github.com/graphql-go/graphql"
)

// GetRootFields declares the query functions.
func GetRootFields() graphql.Fields {
	return graphql.Fields{

		"user":       GetUserQuery(),
		"nearbyUser": GetNearbyUser(),
		"driverList": GetAllDrivers(),
	}
}
