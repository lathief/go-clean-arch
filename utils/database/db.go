package database

import (
	"fmt"
	"log"
	"sesi1-crud/config"
	"sesi1-crud/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%v)/%s?parseTime=True&loc=UTC`, config.Config.Database.DBUser,
		config.Config.Database.DBPass, config.Config.Database.DBHost, config.Config.Database.DBPort,
		config.Config.Database.DBName)
	// dsn := fmt.Sprintf(`root:@tcp(localhost:3306)/crud_go?parseTime=True&loc=UTC`)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connected database ", err)
	}

	err = db.Debug().AutoMigrate(entity.Users{})
	if err != nil {
		log.Fatal("Cannot auto migrate database ", err)
	}
	return db
}
