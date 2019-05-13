package gorm

import (
	"time"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User02 struct {
	gorm.Model
	UserName 	string
	FirstName 	string
	LastName 	string
	Calendar 	Calendar // One-to-One
}

type Calendar struct {
	gorm.Model
	Name 	string
	UserID 	uint // One-to-One
	Appointment []Appointment01 `gorm:"polymorphic:owner"` // One-to-Many   // Polymorphism
}

type Appointment01 struct {
	gorm.Model
	Subject		string
	Description	string
	StartTime	time.Time
	Length		uint
	// CalendarID	uint // One-to-Many
	OwnerID 	uint	// Polymorphism
	OwnerType	string	// Polymorphism
	Attendess	[]User02	`gorm:"many2many:appointment_user"` // Many-to-Many
}

// Polymorphism
type TaskList struct {
	gorm.Model
	Appointment []Appointment01 `gorm:"polymorphic:owner"`
}

/*
	1. One-to-One
	2. Foreign Keys Constraints
	3. One-to-Many
	4. Many-to-Many
	5. Polymorphism
*/
func Relationship() {
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
	db.DropTable(&Appointment01{})
	db.CreateTable(&Appointment01{})

	// 1. One-to-One: insert data to User & Calendar Tables
		// db.Debug().Save(&User{
		// 	UserName: "adent",
		// 	Calendar: Calendar{
		// 		Name: "Improbable Events",
		// 	},
		// })
		// Query of 1.
		// u := User{}
		// c := Calendar{}
		// db.First(&u).Related(&c, "calendar")
		// fmt.Println(u)
		// fmt.Println(c)

	// 2. Foreign Keys Constraints: User <=> Calendar
		// db.Debug().Model(&Calendar{}).
		// 	AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE") // field, dest, onDelete, onUpdate
		// db.Save(&User{
		// 	UserName: "Be Chip",
		// 	Calendar: Calendar{
		// 		Name: "Improbable Events",
		// 	},
		// })
	
	// 3. One-to-Many: Calendar <=> Appointment
		// db.Debug().Save(&User{
		// 		UserName: "Be Chip",
		// 		Calendar: Calendar{
		// 			Name: "Improbable Events",
		// 			Appointment: []Appointment{
		// 				{Subject: "Spontaneous Whale Generation"},
		// 				{Subject: "Saved from Vaccuum of Space"},
		// 			},
		// 		},
		// 	})
	
	// 4. Many-to-Many: User <=> Appointment
		// users := []User{
		// 	{UserName: "fprefect"},
		// 	{UserName: "tmacillan"},
		// 	{UserName: "mrobot"},
		// }

		// for i := range users {
		// 	db.Save(&users[i])
		// } 

		// db.Debug().Save(&User{
		// 		UserName: "Be Chip",
		// 		Calendar: Calendar{
		// 			Name: "Improbable Events",
		// 			Appointment: []Appointment{
		// 				{Subject: "Spontaneous Whale Generation", Attendess: users},
		// 				{Subject: "Saved from Vaccuum of Space", Attendess: users},
		// 			},
		// 		},
		// 	})
	
	// 5. Polymorphism
		users := []User02{
			{UserName: "fprefect"},
			{UserName: "tmacillan"},
			{UserName: "mrobot"},
		}

		for i := range users {
			db.Save(&users[i])
		} 

		db.Debug().Save(&User02{
				UserName: "Be Chip",
				Calendar: Calendar{
					Name: "Improbable Events",
					Appointment: []Appointment01{
						{Subject: "Spontaneous Whale Generation", Attendess: users},
						{Subject: "Saved from Vaccuum of Space", Attendess: users},
					},
				},
			})
	
	// 6. Association Method
		// db.Model(&Calendar{}).Association("Appointments")
		// 					.Find(&appointments)
		// 					.Append(&appointmentsToAdd)
		// 					.Delete(&appointmentsToDelete)
		// 					.Replace(&appointmentsToSubstitute)
		// 					.Count()
		// 					.Clear()

	// FINISH
	fmt.Println("===========> DONE.")
}