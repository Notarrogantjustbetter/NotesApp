package main

import (
	"net/http"

	"github.com/Notarrogantjustbetter/NotesApp/v2/database"
	"github.com/Notarrogantjustbetter/NotesApp/v2/routes"
	"github.com/Notarrogantjustbetter/NotesApp/v2/utils"
)

func main() {
	utils.LoadTemplate()
	database.InitRedis()
	router := routes.InitRouter()
	http.ListenAndServe(":8080", router)
}
