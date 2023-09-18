package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/path/", app.handleLinkRedirection)

	return mux

}
