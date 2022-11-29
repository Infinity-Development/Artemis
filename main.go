package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	mainDomain = "https://infinitybots.gg"
)

type Handler = func(r *http.Request) string

func wrap(fn Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var redirPath = fn(r)

		http.Redirect(w, r, redirPath, http.StatusTemporaryRedirect)
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.CleanPath)

	r.NotFound(wrap(func(r *http.Request) string {
		return mainDomain + r.URL.EscapedPath()
	}))

	// handle / route
	r.Handle("/", wrap(func(r *http.Request) string {
		return mainDomain
	}))

	// We redirect bots onto itself to allow for better redirects
	r.Handle("/bot/{id}", wrap(func(r *http.Request) string {
		return "https://" + r.Host + "/" + chi.URLParam(r, "id")
	}))

	r.Handle("/bot/{id}/{path}", wrap(func(r *http.Request) string {
		return "https://" + r.Host + "/" + chi.URLParam(r, "id") + "/" + chi.URLParam(r, "path")
	}))

	// Bot redirects
	r.Handle("/{id}", wrap(func(r *http.Request) string {
		return mainDomain + "/bots/" + chi.URLParam(r, "id")
	}))

	r.Handle("/{id}/{path}", wrap(func(r *http.Request) string {
		path := chi.URLParam(r, "path")

		var redirPath string

		switch path {
		case "i":
			redirPath = "invite"
		case "inv":
			redirPath = "invite"
		case "invite":
			redirPath = "invite"
		case "v":
			redirPath = "vote"
		case "vote":
			redirPath = "vote"
		default:
			redirPath = path
		}

		return mainDomain + "/bots/" + chi.URLParam(r, "id") + "/" + redirPath
	}))

	// Packs
	for _, p := range []string{
		"/pack/{id}",
		"/packs/{id}",
		"/p/{id}",
		"/pack/{id}/",
		"/packs/{id}/",
		"/p/{id}/",
	} {
		r.Handle(p, wrap(func(r *http.Request) string {
			return mainDomain + "/packs/" + chi.URLParam(r, "id")
		}))
	}

	// User routes
	for _, p := range []string{
		"/u/{id}",
		"/user/{id}",
		"/users/{id}",
		"/profile/{id}",
		"/profiles/{id}",
		"/u/{id}/",
		"/user/{id}/",
		"/users/{id}/",
		"/profile/{id}/",
		"/profiles/{id}/",
	} {
		r.Handle(p, wrap(func(r *http.Request) string {
			return mainDomain + "/users/" + chi.URLParam(r, "id")
		}))
	}

	err := http.ListenAndServe(":1010", r)

	if err != nil {
		panic(err)
	}
}
