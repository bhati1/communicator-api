package router

import (
	"comm-api/controller"

	"github.com/gorilla/mux"
)

// type Item struct {
// 	Message string `json:"message"`
// }

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/get", controller.GetAllMessage).Methods("GET")
	router.HandleFunc("/get/{id}", controller.GetMessageById).Methods("GET")
	router.HandleFunc("/delete", controller.DeleteAllMessages).Methods("DELETE")
	router.HandleFunc("/delete/{id}", controller.DeleteOneMessage).Methods("DELETE")
	router.HandleFunc("/send", controller.SendMessage).Methods("POST")
	router.HandleFunc("/update/{id}/{msg}", controller.UpdateMessageById).Methods("PUT")
	router.HandleFunc("/downloadmsgasfile", controller.DownloadMsgAsFile).Methods("GET")
	router.HandleFunc("/downloadvideo", controller.DownloadVideo).Methods("GET")
	return router
}
