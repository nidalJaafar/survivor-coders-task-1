package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"student/model"
)

var Db *gorm.DB

func DbConfig() {
	dsn := "host=localhost user=postgres password=password dbname=students port=5432"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	err = Db.AutoMigrate(&model.Student{})
	if err != nil {
		return
	}

}
