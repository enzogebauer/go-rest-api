package routes

import (
	"log"
	"net/http"

	"github.com/enzogebauer/go-rest-api.git/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
