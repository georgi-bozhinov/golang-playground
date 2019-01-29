package notes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type NoteController struct {
	Repository *NoteRepository
}

// NewNoteController creates a new note controller with a provided repository
func NewNoteController(repository NoteRepository) *NoteController {
	return &NoteController{Repository: &repository}
}

// GetNotes fetches all notes from the `notes.json` file
func (nc NoteController) GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Will be read from db
	res := nc.Repository.GetNotes()

	if res.Error != nil {
		log.Fatal(res.Error)
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
	}

	notes := res.Notes
	err := json.NewEncoder(w).Encode(&notes)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetNotesById returns a note by its id
func (nc NoteController) GetNoteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	res := nc.Repository.GetNoteById(params["id"])

	if res.Error != nil {
		log.Fatal(res.Error)
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
	}

	note := res.Notes[0]
	err := json.NewEncoder(w).Encode(&note)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
