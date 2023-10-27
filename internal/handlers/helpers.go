package handlers

import (
	"grupie-tracker/internal/parser"
	"net/http"
	"strconv"
	"text/template"
)

type Aplication struct {
	Models []parser.Artists
}

func (app *Aplication) Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.mainPage)
	mux.HandleFunc("/artist", app.artistPage)
	styles := http.FileServer(http.Dir("./internal/web/ui/"))
	mux.Handle("/static/", http.StripPrefix("/static/", styles))
	return mux
}

func (app *Aplication) errors(w http.ResponseWriter, problem int) {
	tmlp := template.Must(template.ParseFiles("./internal/web/html/error.html"))
	e := "problem is " + strconv.Itoa(problem)
	tmlp.Execute(w, e)

}
