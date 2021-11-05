package devicemngmt

import (
	"errors"
	"fmt"

	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBConfig ...
type MongoDBConfig struct {
	Host, User, Password, DBName, mechanism, source string
}

// Config ...
type Config struct {
	// MongoDB config, for save documents
	MongoDB MongoDBConfig
	// Table prefix, each service has its own prefix
	TablePrefix string
}

// Service ...
type Service struct {
	Config
	DB *mongo.Database
}

var s *Service

// Init ...
func Init(config Config) (*Service, error) {
	if config.MongoDB.Host == "" || config.TablePrefix == "" {
		return nil, errors.New("please provide all necessary information for init device")
	}

	// Connect MongoDB
	db, err := mongodb.Connect(
		config.MongoDB.Host,
		config.MongoDB.User,
		config.MongoDB.Password,
		config.MongoDB.DBName,
		config.MongoDB.mechanism,
		config.MongoDB.source,
	)
	if err != nil {
		fmt.Println("Cannot init module DEVICE MANAGEMENT", err)
		return nil, err
	}

	s = &Service{
		Config: config,
		DB:     db,
	}

	return s, nil
}

// GetInstance ...
func GetInstance() *Service {
	return s
}
