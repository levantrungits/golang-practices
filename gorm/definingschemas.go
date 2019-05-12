package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User01 struct {
	// gorm.Model								// auto generate fields GORM default
	Model		gorm.Model `gorm:"embedded"`	// Embedding Child Objects
	UserID		int		`gorm:"primary_key"`
	UserName 	string	`sql:"type:VARCHAR(15); not null"` 				   // Customizing Field Type and Sizes
	FirstName 	string	`sql:"size:100; not null" gorm:"column:FirstName"` // Controlling Column Names
	LastName 	string	`sql:"unique; unique_index; not null; DEFAULT:'CHIP'" gorm:"column:LastName"` // Controlling Column Names
	Count		int 	`gorm:"AUTO_INCREMENT"`  // Auto-incrementing Field ==> ~ ID
	TempField 	bool 	`sql:"-"`				 // Transient Fields ==> not create in db field
}

var users []User = []User {
	User {UserName: "TrungLv", FirstName: "LeVan", LastName: "Trung"},
	User {UserName: "adent", FirstName: "Ford", LastName: "Prefect"},
	User {UserName: "tmacmillan", FirstName: "Tricia", LastName: "MacMillan"},
	User {UserName: "mrobot", FirstName: "Marvin", LastName: "Robot"},
}

/*
	1. `sql:"type:VARCHAR(15)"`				Specify Type
	2. `sql:"size:100"`						Field Size
	3. `gorm:"AUTO_INCREMENT"`				Auto Increment
	4. `sql:"-"`							Ignore a Field
	5. `sql:"unique_index"`					Add a unique index for field
	6. `sql:"not null"`						Set field to require values
	7. `sql:"unique"`						Set field to prevent duplicates
	8. `sql:"DEFAULT:'TrungLv'"`			Set field's default value
	9. `gorm:"primary_key"`					Set field to be primary key
	10. `gorm:"column:New Column Name"`		Set field's name in database
	11. `gorm:"embedded"`					Inline struct owning table
*/
func DefiningSchemas() {
	// CONNECT POSTGRES DB
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=default password=secret dbname=gorm_db sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// INIT TABLE
	db.DropTable(&User01{})
	db.CreateTable(&User01{})

	// ADRANGE USERs MODEL
	for _, user := range users {
		db.Create(&user)
	}

	/* SELECT ONE
	u := User{}
	db.First(&u)
	db.Last(&u)
	*/

	/* UPDATE ONE
	u := User{UserName: "tmacmillan"}
	db.Where(&u).First(&u)
	fmt.Println(u)

	u.LastName = "Bao Tran"
	db.Save(&u)

	user := User{}
	db.Where(&u).First(&user)
	fmt.Println(user)
	*/

	/* DELETE ONE
	db.Where(&User{UserName: "adent"}).Delete(&User{})
	*/

	// Embedding Child Objects - Print All Fields of User Model
	for _, f := range db.NewScope(&User01{}).Fields() {
		fmt.Println(f.Name)
	}

	// Working with Indexes
	db.Model(&User01{}).AddIndex("idx_first_name", "FirstName")
	db.Model(&User01{}).AddUniqueIndex("idx_last_name", "LastName")
	db.Model(&User01{}).RemoveIndex("uix_sys_users_lastname")

	// FINISH
	fmt.Println("=====> Done.")
}

// Define Table Name for User Model
func (u User01) TableName() string {
	return "SYS_Users"
}


