package main

import (
	"log"

	"awacs_smart_api_service/app"
	db "github.com/brkelkar/common_utils/databases"

	cr "github.com/brkelkar/common_utils/configreader"
	"github.com/brkelkar/common_utils/logger"

	"gorm.io/gorm"
)

var (
	envVerables = []string{"SERVER_PORT", "SERVER_HOST", "AWACS_DB", "AWACS_SMART_DB", "AWACS_SMART_STOCKIST_DB",
		"DB_PORT", "DB_HOST", "DB_USERNAME", "DB_PASSWORD", "GRPC_SERVER", "GRPC_CLIENT"}
	err error
)
func main() {
	logger.Info("Starting Sync upload service")

	var cfg cr.Config
	//Read configuration from config.yml file
	cfg.ReadFile("config.yml")

	//Read enviroment variables
	m := cfg.ReadEnv(envVerables)

	//Over write config file variables if enviroment variable is set
	cfg.MapEnv(m)

	// Create connention map incase of multiple db
	// map will hold value for *db with respective string identifier
	var dbObj db.DbObj

	db.DB = make(map[string]*gorm.DB, 3)
	db.DB["smartdb"], err = dbObj.GetConnection("smartdb", cfg)
	if err != nil {
		logger.Error("Error while connecting to smartdb", err)
		log.Fatal(err)
	}

	dbObj.CreateConnectionPool(db.DB["smartdb"], 10, 100, 60)
	logger.Info("Connection successful smartdb")

	db.DB["awacs"], err = dbObj.GetConnection("awacs", cfg)
	if err != nil {
		logger.Error("Error while connecting to awacs", err)
		log.Fatal(err)
	}

	dbObj.CreateConnectionPool(db.DB["awacs"], 10, 100, 60)
	logger.Info("Connection successful awacs")
	app.Start_Application(cfg)

}
