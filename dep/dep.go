package dep

import (
	"log"

	"github.com/go-bongo/bongo"
	"github.com/personal-website-go/conf"
	"github.com/personal-website-go/dao"
	"github.com/personal-website-go/service"
)

// Local dependencies
var connection *bongo.Connection
var contentDao dao.ContentDao

// Global dependencies
var ContentService service.ContentService

func initDatabase() {
	config := &bongo.Config{}
	config.ConnectionString = conf.Config().DatabaseConfig.ConnectionString
	config.Database = conf.Config().DatabaseConfig.DatabaseName
	var err error
	connection, err = bongo.Connect(config)
	if err != nil {
		log.Fatalf("Unable to connect to MongoDB : %v", err.Error())
	}
}

func InitDep() {
	initDatabase()
	contentDao.Init(connection)
	ContentService.Init(contentDao)
}
