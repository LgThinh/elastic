package service

import (
	"es-tranform/conf"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func StartES(address string) (es *elasticsearch.Client) {
	//cfg := elasticsearch.Config{
	//	Addresses: []string{
	//		address,
	//	},
	//}
	//es, err := elasticsearch.NewClient(cfg)
	//if err != nil {
	//	log.Fatal("Error creating client:%s", err)
	//}
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response ES", err)
	}
	defer res.Body.Close()
	return es
}

func ConnectDB(conf conf.AppConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", conf.DBHost, conf.DBUser, conf.DBPass, conf.DBName, conf.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
