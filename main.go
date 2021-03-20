package main

import (
	"net/http"
	"github.com/Notarrogantjustbetter/NotesApp/v2/database"
	"github.com/Notarrogantjustbetter/NotesApp/v2/server"
	"github.com/Notarrogantjustbetter/NotesApp/v2/utils"
)


func main() {
	utils.LoadTemplate()
	database.InitRedis()
	router := server.InitServer()
	http.ListenAndServe(":8080", router)
}