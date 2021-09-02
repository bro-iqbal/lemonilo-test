package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"lemonilo/helpers"
	"lemonilo/models"

	"github.com/gorilla/mux"
	"github.com/thedevsaddam/govalidator"
)

func UserList(rw http.ResponseWriter, r *http.Request) {
	var users models.Users
	var email, address, search string

	if r.URL.Query()["Email"] != nil {
		email = r.URL.Query()["Email"][0]
	}
	if r.URL.Query()["Address"] != nil {
		address = r.URL.Query()["Address"][0]
	}
	if r.URL.Query()["Search"] != nil {
		search = r.URL.Query()["Search"][0]
	}

	var query = ""

	if email != "" {
		if query != "" {
			query = query + " AND "
		}
		query = query + "email = " + fmt.Sprintf("'%v'", email)
	} else if address != "" {
		if query != "" {
			query = query + " AND "
		}
		query = query + "address = " + fmt.Sprintf("'%v'", address)
	} else if search != "" {
		if query != "" {
			query = query + " AND "
		}
		query = query + "email LIKE '%" + fmt.Sprintf("%v", search) + "%' " + "OR address LIKE '%" + fmt.Sprintf("%v", search) + "%'"
	}

	result := db.Where(query).Find(&users)

	if result.Error != nil {
		helpers.HandleResponse(rw, 400, "Get user failed", result.Error)

		return
	}

	helpers.HandleResponse(rw, 200, "Get user list successful", users)

	return
}

func UserDetail(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	result := db.Where("userid = ?", id).First(&user)

	if result.Error != nil {
		helpers.HandleResponse(rw, 400, "Get user by ID "+fmt.Sprintf("%v", id)+" failed", result.Error)

		return
	}

	helpers.HandleResponse(rw, 200, "Get user by ID "+fmt.Sprintf("%v", id)+" successful", user)

	return
}

func UserCreate(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	validator := govalidator.New(govalidator.Options{
		Request: r,
		Rules: govalidator.MapData{
			"Email":    []string{"required: The Email field is required"},
			"Username": []string{"required: The Username field is required"},
			"Address":  []string{"required: The Address field is required"},
			"Password": []string{"required: The Password field is required"},
		},
		RequiredDefault: true,
	}).Validate()

	if len(validator) > 0 {
		helpers.HandleResponse(rw, 400, "Validation error", validator)

		return
	}

	pass, _ := helpers.HashPassword(r.FormValue("Password"))

	user.Email = r.FormValue("Email")
	user.Username = r.FormValue("Username")
	user.Address = r.FormValue("Address")
	user.Password = pass

	result := db.Create(&user)

	if result.Error != nil {
		helpers.HandleResponse(rw, 400, "Create user failed", result.Error)

		return
	}

	db.Order("userid DESC").First(&user)

	helpers.HandleResponse(rw, 200, "Create user successful", user)

	return
}

func UserUpdate(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	validator := govalidator.New(govalidator.Options{
		Request: r,
		Rules: govalidator.MapData{
			"Email":    []string{},
			"Username": []string{},
			"Address":  []string{},
		},
		RequiredDefault: true,
	}).Validate()

	if len(validator) > 0 {
		helpers.HandleResponse(rw, 400, "Validation error", validator)

		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user.UserID = id
	user.Email = r.FormValue("Email")
	user.Username = r.FormValue("Username")
	user.Address = r.FormValue("Address")

	result := db.Where("userid = ?", id).Model(&user).Updates(user)

	if result.Error != nil {
		helpers.HandleResponse(rw, 400, "Update user id "+string(id)+" failed", result.Error)

		return
	}

	helpers.HandleResponse(rw, 200, "Update user successful", user)

	return
}

func UserUpdatePassword(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	validator := govalidator.New(govalidator.Options{
		Request: r,
		Rules: govalidator.MapData{
			"OldPassword":     []string{"required: The Old Password field is required"},
			"ConfirmPassword": []string{"required: The Confirm Password field is required"},
			"NewPassword":     []string{"required: The New Password field is required"},
		},
		RequiredDefault: true,
	}).Validate()

	if len(validator) > 0 {
		helpers.HandleResponse(rw, 400, "Validation error", validator)

		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	db.Where("userid = ?", id).First(&user)

	if r.FormValue("OldPassword") == r.FormValue("confirmPassword") {
		match := helpers.CheckPasswordHash(r.FormValue("OldPassword"), user.Password)
		if !match {
			helpers.HandleResponse(rw, 400, "Update user password "+string(id)+" failed", nil)

			return
		}
	}

	pass, _ := helpers.HashPassword(r.FormValue("NewPassword"))

	user.Password = pass

	result := db.Where("userid = ?", id).Model(&user).Updates(user)

	if result.Error != nil {
		helpers.HandleResponse(rw, 400, "Update user password "+string(id)+" failed", result.Error)

		return
	}

	helpers.HandleResponse(rw, 200, "Update user password successful", result.Value)

	return
}

func UserDelete(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	result := db.Model(&user).Where("userid = ?", id).Delete(user)

	if result.Error != nil {
		helpers.HandleResponse(rw, 400, "Delete user id "+string(id)+" failed", result)

		return
	}

	helpers.HandleResponse(rw, 200, "Delete user successful", map[string]interface{}{
		"ID": id,
	})

	return
}
