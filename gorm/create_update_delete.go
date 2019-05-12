package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Create_Update_Delete() {
	// CONNECT POSTGRES DB
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=default password=secret dbname=gorm_db sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// INIT TABLE
	db.DropTable(&User{})
	db.CreateTable(&User{})
	db.DropTable(&Calendar{})
	db.CreateTable(&Calendar{})
	db.DropTable(&Appointment{})
	db.CreateTable(&Appointment{})


	fmt.Println("===========> DONE.")
}