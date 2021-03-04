package config

import (
	redis "github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
	"log"
)

// CreateRedisClient creates a new redis client instance.
func CreateRedisClient() *redis.Client{

	// Set the file name of the configurations file
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")

	viper.SetConfigType("yml")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()


	var configuration Configurations


	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)

	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	return redis.NewClient(&redis.Options{
	Addr: configuration.Redis.Addr,
	}) 
}