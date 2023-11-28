package handlers

import (
	"encoding/json"
	"fmt"
	"grupie-tracker/internal/parser"
	"html/template"
	"index/suffixarray"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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

	tmlp, err := template.ParseFiles("./internal/web/html/index.html")
	if err != nil {
		app.errors(w, http.StatusBadGateway)
		return
	}
	err = tmlp.ExecuteTemplate(w, "index", app.Models)
	if err != nil {
		app.errors(w, http.StatusInternalServerError)
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

func (app *Aplication) search(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		return
	}

	inp := r.URL.Query().Get("text")

	var matchingModels []parser.Artists

	for _, x := range app.Models {
		if strings.Contains(inp, x.Name) {
			matchingModels = append(matchingModels, x)
		}
		if containArray(inp, x.Members) {
			matchingModels = append(matchingModels, x)
		}
	}

	if len(matchingModels) == 0 {
		// http.Error(w, "Not Found", http.StatusNotFound)
	}

	jsonData, err := json.Marshal(matchingModels)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

func containArray(str string, subStrs []string) bool {
	if len(subStrs) == 0 {
		return true
	}
	r := regexp.MustCompile(strings.Join(subStrs, "|"))
	index := suffixarray.New([]byte(str))
	res := index.FindAllIndex(r, -1)
	exists := make(map[string]int)
	for _, v := range subStrs {
		exists[v] = 1
	}
	for _, pair := range res {
		s := str[pair[0]:pair[1]]
		exists[s] = exists[s] + 1
	}
	for _, v := range exists {
		if v == 1 {
			return false
		}
	}
	return true
}
