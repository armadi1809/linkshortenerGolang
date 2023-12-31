package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/path/", app.handleLinkRedirection)
	mux.HandleFunc("/", app.homePage)
	mux.HandleFunc("/shorten", app.linkResponse)

	return mux

}
