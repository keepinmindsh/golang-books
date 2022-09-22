package tips

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type HelloHandler struct {
	db *sql.DB
}

func (h *HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var name string
	// Execute the query.
	row := h.db.QueryRow("SELECT myname FROM mytable")
	if err := row.Scan(&name); err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	// Write it back to the client.
	fmt.Fprintf(writer, "hi %s!\n", name)
}

func HttpHandlerStart() {
	// Open our database connection.
	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatal(err)
	}
	// Register our handler.
	http.Handle("/hello", &HelloHandler{db: db})
	http.ListenAndServe(":8080", nil)

}
