package controllers

import (
	"lemonilo/database"
	"lemonilo/helpers"
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

var JWTKeys = []interface{}{"iGSpuiPe6YubU7TRu0NZ8dHBUutMIXZvn194MHyfhFwQJ4VTNYNn1qErusfNdgTD", "r0EPsjAcrho8VHeSO2MYinhrnbFwgZIb"}

func Init() {
	db = database.Connection()

	return
}

func Documentation(rw http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"documentation": "https://documenter.getpostman.com/view/1585979/U16dT9KH",
	}

	helpers.HandleResponse(rw, 200, "Update user successful", data)

	return

}
