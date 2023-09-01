package migrations

import (
	"go-bank-backend/helpers"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialect/postgres"
	"gorm.io/gorm"
	"honnef.co/go/tools/analysis/facts/generated"
)

type User struct {
	 gorm.Model
	 Email string `gorm:"type:varchar(100);unique_index"`
	 Password string `gorm:"type:varchar(100)"`
	 Username string `gorm:"type:varchar(100)"`
}

type Account struct {
	gorm.Model
	Type string
	Name string
	Balance uint
	UserID uint

}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=2002 user=postgres dbname=bank password=postgres sslmode=disable")
	helpers.HandleErr(err)
	return db
}

func CreateAccounts() {
	db := connectDB()

	users :=[2]User{
		{Username: "Allan" , Email:"allanmuharikamau@gmail.com"},
		{Username: "Ezra" , Email:"Ezrakanake@gmail.com"},
	}
	for i := 0; i< len(users); i++{
		generatedPassword := helpers.HashandSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := Account{Type: "Daily Account", Name:string(users[i].Username + "s" + "account"), Balance: uint(1000* int(i+1)), UserID:user.ID}
		db.Create(&account)
	
}
defer db.Close()
}

func Migrate() {}