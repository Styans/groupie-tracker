package service

import (
	"fmt"
	"grupie-tracker/internal/handlers"
	"grupie-tracker/internal/parser"
	"net/http"
	"os"
)

func RunServ() {
	model, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR in %v", err)
		return
	}
	app := &handlers.Aplication{
		Models: model,
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.Route(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR in %v", err)
		return
	}
}
