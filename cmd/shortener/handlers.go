package main

import (
	"net/http"
	"strings"
)

func (app *application) handleLinkRedirection(w http.ResponseWriter, r *http.Request) {

	var path string
	if strings.HasPrefix(r.URL.Path, "/path/") {
		path = r.URL.Path[len("/path/"):]
	}

	url, err := app.urls.GetUrl(path)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url.Url, http.StatusPermanentRedirect)
	return
}
