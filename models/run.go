package models

import (
	"fmt"
	"net/http"
	"time"
)

func IndexCreatDev(w http.ResponseWriter, r *http.Request) {
	t, err := time.Parse(time.RFC3339, "2018-04-06T10:49:05Z")
	if err != nil {
		fmt.Println("dome")
	}
	w.Write([]byte())
}
