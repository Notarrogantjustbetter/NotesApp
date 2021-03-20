package server

import (
	"fmt"
	"net/http"
	"github.com/Notarrogantjustbetter/NotesApp/v2/database"
	"github.com/Notarrogantjustbetter/NotesApp/v2/middleware"
	"github.com/Notarrogantjustbetter/NotesApp/v2/sessions"
	"github.com/Notarrogantjustbetter/NotesApp/v2/utils"
	"github.com/gorilla/mux"
)


type Server struct {
	router *mux.Router
}

func InitServer() *mux.Router {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.routes()
	return s.router
}

func (s Server) routes() {
	s.router.HandleFunc("/", middleware.MiddlewareAuthentication(homeHandler().ServeHTTP)).Methods("GET")
	s.router.HandleFunc("/register", getRegisterHandler().ServeHTTP).Methods("GET")
	s.router.HandleFunc("/login", getLoginHandler().ServeHTTP).Methods("GET")
	s.router.HandleFunc("/register", postRegisterHandler().ServeHTTP).Methods("POST")
	s.router.HandleFunc("/login", postLoginHandler().ServeHTTP).Methods("POST")
	s.router.HandleFunc("/", addNoteHandler().ServeHTTP).Methods("POST")
	s.router.HandleFunc("/Notes", getNotesHandler().ServeHTTP).Methods("GET")
	s.router.HandleFunc("/deleteNotes", deleteNoteGetHandler().ServeHTTP).Methods("GET")
	s.router.HandleFunc("/deleteNotes", deleteNotePostHandler().ServeHTTP).Methods("POST")
}

func homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ExecuteTemplate(w, "home.html", nil)
	}
}

func getRegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ExecuteTemplate(w, "register.html", nil)
	}
}

func getLoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ExecuteTemplate(w, "login.html", nil)
	}
}

func postRegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.PostForm.Get("Username")
		password := r.PostForm.Get("Password")
		register := database.User{}.RegisterUser(username, password)
		if register != nil {
			fmt.Fprint(w, "Failed to register.")
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func postLoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.PostForm.Get("Username")
		password := r.PostForm.Get("Password")
		login := database.User{}.LoginUser(username, password)
		if login != nil {
			fmt.Fprint(w, "Failed to login")
		}
		sessions.SetSession(w, r, username)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func addNoteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		note := r.PostForm.Get("Note")
		addNote := database.Note{}.AddNote(note)
		if addNote != nil {
			fmt.Fprint(w, "Failed to add note")
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func getNotesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context, _ := database.Note{}.GetNotes()
		utils.ExecuteTemplate(w, "notes.html", context)
	}
}

func deleteNoteGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ExecuteTemplate(w, "deleteNote.html", nil)
	}
}

func deleteNotePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		note := r.PostForm.Get("Note")
		deleteNote := database.Note{}.DeleteNote(note)
		if deleteNote != nil {
			fmt.Fprint(w, "Failed to delete note.")
		}
		http.Redirect(w, r, "/Notes", http.StatusFound)
	}
}
