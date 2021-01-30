package migration

import (
	"Humo1/helpers"
	"Humo1/interface"
)


func createAccount()  {
	db := helpers.ConnectDB()

	users := &[2]_interface.User{
		{Username: "Jones", Email: "jones@mail.com"},
		{Username: "Brojan", Email: "brojan@mail.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &_interface.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &_interface.Account{Type: "Daily Account", Name: string(users[i].Username + "'s"+ "account"), Balance: uint(
			1000 * int(i + 1))}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate()  {
	db := helpers.ConnectDB()
	db.AutoMigrate(&_interface.User{}, &_interface.Account{})
	defer db.Close()

	createAccount()
}