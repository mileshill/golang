package app

import (
	"github.com/miles1618/udemy/rest/mvc/controllers"
	"log"
	"net/http"
)

func StartApp() {
	// Routes
	http.HandleFunc("/users", controllers.GetUser)
	// Launch server
	err:= http.ListenAndServe(":8050", nil)
	if err != nil {
		log.Fatal(err)
	}
}
