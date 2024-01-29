package config

import (
	"notes_go/handler"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func InitHandler() {
	Router.HandleFunc("/notes", handler.GetAllNotesHandler).Methods("Get")
	Router.HandleFunc("/notes", handler.CreateNoteHandler).Methods("Post")
	Router.HandleFunc("/notes/{id}", handler.GetNoteByIDHandler).Methods("Get")
	Router.HandleFunc("/notes", handler.ModifyNoteHandler).Methods("Put")
	Router.HandleFunc("/notes/{id}", handler.RemoveNoteByIDHandler).Methods("Delete")
}
