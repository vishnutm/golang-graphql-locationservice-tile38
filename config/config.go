package config

// Configurations is the structure for all default/constant values in this microservice
type Configurations struct {
	Server ServerConfigurations
	Redis  RedisConfigurations
}

// ServerConfigurations is the structure for all default/constant values related to the server
type ServerConfigurations struct {
	Port string `yaml:"port"`
}

// RedisConfigurations is the structure for all default/constant values related to the go-redis client
type RedisConfigurations struct {
	Addr string `yaml:"addr"`
}

// Client is the redis-client instance used to communicate with the tile38 server
var Client = CreateRedisClient()
