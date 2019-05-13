package gorm

import (
	"time"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	FirstName 	string
	LastName 	string
	Appointments	[]Appointment // with Children
}
// Define Table Name
func (u User) TableName() string {
	return "SYS_Users"
}

// with Children
type Appointment struct {
	gorm.Model
	UserID uint
	StartTime *time.Time
	Duration uint
	Attendess []*User
	Subject string
	Description string
}

/*
	1. Creating Records
	2. Updating Records
	3. Deleting Records
	4. Transactions
*/
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
	db.DropTable(&Appointment{})
	db.CreateTable(&Appointment{})

	// 1. Creating Records
		u := User{
			FirstName: "Trung",
			LastName: "Le Van",
		}
		// 1.1 Creating Records with Children
		childs := []Appointment{
			Appointment{Subject: "First"},
			Appointment{Subject: "Second", Attendess: []*User{&u}},
			Appointment{Subject: "Third"},
		}
		u.Appointments = childs
		db.Create(&u) // ~ db.Debug().Create(&u)
		//fmt.Println(db.NewRecord(&u)) // False => constraints
	
	// 2. Updating Records
		u.LastName = "ChipChip"
		u.FirstName = "Google"
		db.Save(&u) // ~ db.Debug().Save(&u)
		
		db.Model(&u).Update("first_name", "Facebook")

		db.Model(&u).Updates(
			map[string]interface{}{
				"first_name": "Skype",
				"last_name": "Slack",
			},
		)
		// 2.2 Updating Records without Callbacks

		fmt.Println(u)

	fmt.Println("===========> DONE.")
}