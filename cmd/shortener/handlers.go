package main

import (
	"fmt"
	"html/template"
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
	http.Redirect(w, r, fmt.Sprintf("https://www.%s", url.Path), http.StatusPermanentRedirect)
	return
}

func (app *application) homePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./ui/html/base.html"))
	t.ExecuteTemplate(w, "base", nil)

	return
}

func (app *application) linkResponse(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	formData := r.Form
	url := formData.Get("url")
	if url == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortenedUrl, err := app.urls.CreatePath(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("The short url for %s is %s", shortenedUrl.Path, shortenedUrl.Url)))

	return
}
