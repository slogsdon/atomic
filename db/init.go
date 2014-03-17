package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/slogsdon/atomic/models"
)

var DB gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "user=jefivorgjpbczf password=NmOXHlxsLOGzZaQfSRMSjgYy-5 host=ec2-54-204-47-70.compute-1.amazonaws.com port=5432 dbname=d29t79smbgiahh sslmode=require")

	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	DB.AutoMigrate(models.Prompt{})
	DB.AutoMigrate(models.Response{})
}
