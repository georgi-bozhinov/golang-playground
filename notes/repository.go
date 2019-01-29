package notes

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type NoteRepository struct {
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	conn       *sql.DB
}

type Result struct {
	Notes []Note
	Error error
}

func NewRepository(host, name, user, password string) *NoteRepository {
	n := &NoteRepository{DBHost: host, DBName: name, DBPassword: password, DBUser: user}
	db, err := sql.Open("postgres", n.buildQueryString())

	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database! Error: %s", err))
	}

	n.conn = db

	return n
}

// Connect connects to the db
func (n *NoteRepository) buildQueryString() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", n.DBUser, n.DBPassword, n.DBHost, n.DBName)
}

func (n *NoteRepository) GetNotes() Result {
	dbQuery := "SELECT * FROM notes"
	rows, err := n.conn.Query(dbQuery)
	if err != nil {
		return Result{nil, err}
	}

	var notes []Note
	for rows.Next() {
		var note Note

		if err := rows.Scan(&note.ID, &note.Title, &note.Description); err != nil {
			return Result{nil, err}
		}
		notes = append(notes, note)
	}

	return Result{notes, nil}
}

func (n *NoteRepository) GetNoteById(id string) Result {
	rows, err := n.conn.Query("SELECT * FROM notes WHERE id = $1", id)

	if err != nil {
		return Result{nil, err}
	}

	var note Note
	for rows.Next() {
		if err := rows.Scan(&note.ID, &note.Title, &note.Description); err != nil {
			return Result{nil, err}
		}
	}

	return Result{[]Note{note}, nil}
}
