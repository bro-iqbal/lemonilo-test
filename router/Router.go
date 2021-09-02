package router

import (
	"log"
	"net/http"
	"os"

	"lemonilo/controllers"

	"lemonilo/helpers"

	"lemonilo/middleware"

	"github.com/gorilla/mux"
)

func Route() {
	router := mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		helpers.HandleResponse(rw, 404, "Route not found", nil)

		return
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		helpers.HandleResponse(rw, 405, "Method not allowed", nil)

		return
	})

	v1 := "/api/v1/"

	router.HandleFunc("/documentation", controllers.Documentation).Methods("GET")

	router.HandleFunc(v1+"login", controllers.Login).Methods("POST")

	router.HandleFunc(v1+"users", middleware.Auth(controllers.UserList)).Methods("GET")
	router.HandleFunc(v1+"user/{id}", middleware.Auth(controllers.UserDetail)).Methods("GET")
	router.HandleFunc(v1+"user", middleware.Auth(controllers.UserCreate)).Methods("POST")
	router.HandleFunc(v1+"user/{id}", middleware.Auth(controllers.UserUpdate)).Methods("PUT")
	router.HandleFunc(v1+"password/{id}", middleware.Auth(controllers.UserUpdatePassword)).Methods("PUT")
	router.HandleFunc(v1+"user/{id}", middleware.Auth(controllers.UserDelete)).Methods("DELETE")

	appUrl := os.Getenv("APP_URL")

	appPort := os.Getenv("APP_PORT")

	log.Println("[*] Running at " + appUrl + ":" + appPort)

	log.Fatal(http.ListenAndServe(":"+appPort, router))
}
