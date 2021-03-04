package mutations

import (
	"github.com/graphql-go/graphql"
)

// GetRootFields declares the mutation functions
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createUser":  CreateUserMutation(),
		"setLocation": UpdateLocationMutation(),
		
	}
}
