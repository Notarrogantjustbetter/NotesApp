package routes

import (
	"fmt"
	"net/http"

	"github.com/Notarrogantjustbetter/NotesApp/v2/database"
	"github.com/Notarrogantjustbetter/NotesApp/v2/middleware"
	"github.com/Notarrogantjustbetter/NotesApp/v2/sessions"
	"github.com/Notarrogantjustbetter/NotesApp/v2/utils"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.ExecuteTemplate(w, "register.html", nil)
	case "POST":
		r.ParseForm()
		username := r.PostForm.Get("Username")
		password := r.PostForm.Get("Password")
		register := database.RegisterUser(username, password)
		if register != nil {
			fmt.Fprint(w, "Failed to register")
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.ExecuteTemplate(w, "login.html", nil)
	case "POST":
		r.ParseForm()
		username := r.PostForm.Get("Username")
		password := r.PostForm.Get("Password")
		login := database.LoginUser(username, password)
		if login != nil {
			fmt.Fprint(w, "Failed to login")
		}
		sessions.SetSession(w, r, username)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func addNoteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	note := r.PostForm.Get("Note")
	addNote := database.AddNote(note)
	if addNote != nil {
		fmt.Fprint(w, "Failed to add note")
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func getNotesHandler(w http.ResponseWriter, r *http.Request) {
	context, _ := database.GetNotes()
	utils.ExecuteTemplate(w, "notes.html", context)
}

func deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.ExecuteTemplate(w, "deleteNote.html", nil)
	case "POST":
		r.ParseForm()
		note := r.PostForm.Get("Note")
		deleteNote := database.DeleteNote(note)
		if deleteNote != nil {
			fmt.Fprint(w, "Failed to delete note")
		}
		http.Redirect(w, r, "/Notes", http.StatusFound)
	}
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", middleware.MiddlewareAuthentication(homeHandler)).Methods("GET")
	router.HandleFunc("/register", registerHandler).Methods("GET", "POST")
	router.HandleFunc("/login", loginHandler).Methods("GET", "POST")
	router.HandleFunc("/", addNoteHandler).Methods("POST")
	router.HandleFunc("/deleteNotes", deleteNoteHandler).Methods("GET", "POST")
	router.HandleFunc("/Notes", getNotesHandler).Methods("GET")
	return router
}
