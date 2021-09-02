package controllers

import (
	"lemonilo/helpers"
	"lemonilo/models"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	username := r.FormValue("username")
	password := r.FormValue("password")

	result := db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		helpers.HandleResponse(rw, 400, "Login failed", result.Error)

		return
	}

	pwdMatch := helpers.Verify(user.Password, password)

	if username != user.Username || !pwdMatch {
		helpers.HandleResponse(rw, 400, "Login failed", result.Error)

		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.UserID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		helpers.HandleResponse(rw, 400, "Login failed", result.Error)

		return
	}

	data := map[string]interface{}{
		"token": t,
	}

	helpers.HandleResponse(rw, 200, "Login successful", data)

	return
}
