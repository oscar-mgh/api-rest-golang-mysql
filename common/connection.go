package common

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/oscar-mgh/api-rest-golang-mysql/model"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/test?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()
	log.Println("Migrating...")
	db.AutoMigrate(&model.Person{})
}

func Seed() {
	db := GetConnection()
	defer db.Close()
	db.Exec(`INSERT into people(id, first_name, last_name, address, phone_number)
	VALUES (1, 'Alice', 'Comley', '7 Harper Circle', '351-355-6190');`)
	db.Exec(`INSERT into people(id, first_name, last_name, address, phone_number)
	VALUES (2, 'Nanny', 'Goodman', '9752 Magdeline Alley', '609-685-2097');`)
	db.Exec(`INSERT into people(id, first_name, last_name, address, phone_number)
	VALUES (3, 'Tessa', 'Jackson', '0313 Scofield Center', '601-134-7729');`)
	db.Exec(`INSERT into people(id, first_name, last_name, address, phone_number)
	VALUES (4, 'Tuesday', 'McKendo', '28442 Oriole Hill', '617-863-4435');`)
	db.Exec(`INSERT into people(id, first_name, last_name, address, phone_number)
	VALUES (5, 'Wyatt', 'Weich', '49350 Coolidge Lane', '885-283-1481');`)
	db.Exec(`INSERT into people(id, first_name, last_name, address, phone_number)
	VALUES (6, 'Reuben', 'Ledwich', '481 Northfield Pass', '325-429-0174');`)
	db.Exec(`INSERT into people(id, first_name, last_name, address, phone_number)
	VALUES (7, 'Daniel', 'Van Waadenburg', '66138 North Place', '681-602-9248');`)
	db.Exec(`INSERT into people(id, first_name, last_name, address, phone_number)
	VALUES (8, 'Johnny', 'Hazeldine', '7493 Elka Drive', '945-883-1624');`)
	db.Exec(`INSERT into people(id, first_name, last_name, address, phone_number)
	VALUES (9, 'Sarah', 'White', '5593 Fallview Alley', '229-776-5392');`)
	db.Exec(`INSERT into people(id, first_name, last_name, address, phone_number)
	VALUES (10, 'Alexa', 'Halton', '953 Division Drive', '578-389-9640');`)
}
