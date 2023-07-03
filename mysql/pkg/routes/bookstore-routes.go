package routes
import (
	"github.com/gorilla/mux"
	"../controllers"

)
var RegisterBook=func(router *mux.Router){
	router.HandleFunc("/book/",controllers.Createbook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookbyid).Methods("GET")
	router.HandleFunc("/book/{id}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}",controllers.deletebook).Methods("DELETE")

}

