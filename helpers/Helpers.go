package helpers

import (
	"encoding/json"
	"fmt"
	"lemonilo/models"
	"log"
	"net/http"
	"reflect"

	"lemonilo/structs"

	jwt "github.com/dgrijalva/jwt-go"
	echo "github.com/labstack/echo/v4"
	bcrypt "golang.org/x/crypto/bcrypt"
)

func Verify(pass string, reqPass string) bool {
	byteHash := []byte(pass)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(reqPass))

	if err != nil {
		return false
	}

	return true
}

func GetTokenID(c echo.Context) int {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	idToken := claims["id"].(float64)
	id := int(idToken)

	return id
}

func VerifyToken(c echo.Context, user models.User) bool {
	reqToken := c.Request().Header.Get("Authorization")
	token := "Bearer " + user.Token
	if token != reqToken {
		return false
	}

	return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func JSONEncode(data interface{}) string {
	jsonResult, _ := json.Marshal(data)

	return string(jsonResult)
}

func HandleError(message string, err interface{}) {
	log.Println()
	log.Println("========== Start Error Message ==========")
	log.Println("Message => " + message + ".")

	if err != nil {
		log.Println("Error => ", err)
	}

	log.Println("========== End Of Error Message ==========")
	log.Println()
}

func HandleResponse(response http.ResponseWriter, code int, message string, data interface{}) {
	var responseStruct = new(structs.Response)

	if code == 200 || code == 201 || code == 202 {
		responseStruct.Success(code, message, data)
	} else {
		HandleError(message, data)

		if message == "Data not found" {
			code = 404
		}

		if data == nil {
			responseStruct.Error(code, message, nil)
		} else if fmt.Sprintf("%v", reflect.TypeOf(data).Kind()) == "ptr" {
			responseStruct.Error(code, message, fmt.Sprintf("%v", data))
		} else {
			responseStruct.Error(code, message, data)
		}
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)
	response.Write([]byte(JSONEncode(responseStruct)))
}
