package singleton

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	dns := "host=localhost user=admin password=1029384756 dbname=main port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
	return db
}
func GetPasswordKey() string {
	return "dgHigFWosugKStak"
}
