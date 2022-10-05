package devicemngmt

import (
	"errors"
	"fmt"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// Config ...
type Config struct {
	// MongoDB config, for save documents
	MongoDB mongodb.Config
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
	if config.MongoDB.Host == "" {
		return nil, errors.New("please provide all necessary information for init device")
	}

	// If prefixTable is empty then it is devicemngmt
	if config.TablePrefix == "" {
		config.TablePrefix = tablePrefixDefault
	}

	// Connect MongoDB
	db, err := mongodb.Connect(config.MongoDB)
	if err != nil {
		fmt.Println("Cannot init module DEVICE MANAGEMENT", err)
		return nil, err
	}

	logger.Init(fmt.Sprintf("%s-devicemngmt", config.TablePrefix), "")

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

// GetConnectOptions ...
func GetConnectOptions(Host, DBName string) mongodb.Config {
	return mongodb.Config{
		Host:       Host,
		DBName:     DBName,
		Standalone: &mongodb.ConnectStandaloneOpts{},
		TLS:        &mongodb.ConnectTLSOpts{},
	}
}
