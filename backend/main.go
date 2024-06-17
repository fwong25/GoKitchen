package main

import (
	"db_mgmt"
	interfaces "interfaces"
	"net/http"
)

func main() {
	db_mgmt.ConnectDB()

	http.HandleFunc("/", interfaces.Handler)
	http.ListenAndServe(":8000", nil)

	db_mgmt.CloseDB()
}
