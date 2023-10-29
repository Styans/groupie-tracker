package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *Aplication) mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.errors(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.errors(w, http.StatusMethodNotAllowed)
		return
	}

	tmlp := template.Must(template.ParseFiles("./internal/web/html/index.html"))
	err := tmlp.ExecuteTemplate(w, "index", app.Models)
	if err != nil {
		app.errors(w, http.StatusNotFound)
		return
	}
}

func (app *Aplication) artistPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		app.errors(w, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 || id > 52 {
		app.errors(w, http.StatusNotFound)
		return
	}

	tmlp := template.Must(template.ParseFiles("./internal/web/html/artist.html"))
	err = tmlp.ExecuteTemplate(w, "artist", app.Models[id-1])
	if err != nil {
		fmt.Println("as")
		app.errors(w, http.StatusNotFound)
		return
	}

}
