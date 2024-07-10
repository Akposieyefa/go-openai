package internals

import (
	"log"

	"github.com/akposieyefa/open-ai/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error
	DB, err = gorm.Open(mysql.Open(config.Envs.DB_URL), &gorm.Config{})

	if err != nil {
		log.Fatal("Sorry unable to connect to database", err)
		return
	}
}
