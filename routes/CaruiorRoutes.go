package routes

import (
	"github.com/Toanlele/Caruio/models"
	"github.com/gorilla/mux"
)

type carinfo struct {
	uid    uint16 //功能号
	status uint8  //状态
}

var DOMCar = func(router *mux.Router) {
	router.HandleFunc("/dev", models.IndexCreatDev).Methods("POST")
}
