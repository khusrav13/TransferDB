package users

import (
	"Humo1/helpers"
	_interface "Humo1/interface"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Login(username string, pass string) map[string]interface{}  {

	db := helpers.ConnectDB()
	user := &_interface.User{}
	if db.Where("username = ?", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Wrong password"}
	}

	accounts := []_interface.ResponseAccount{}
	db.Table("account").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)


	responseUser := &_interface.ResponseUser{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Account: accounts,
	}

	defer db.Close()


	///Set Token

	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry": time.Now().Add(time.Minute ^ 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString("TokenPassword")
	helpers.HandleErr(err)

	var response = map[string]interface{}{"message": "all is fine"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}