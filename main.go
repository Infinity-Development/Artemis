package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	domainShort = []string{"botlist.site", "127.0.0.1:1010"}
	shortRoutes = make(map[string]func(r *http.Request, vars map[string]string) string)
)

const (
	mainDomain = "https://infinitybots.gg"
	prod       = false
)

func isShort(url string) bool {
	fmt.Println(url)
	for _, v := range domainShort {
		if url == v {
			return true
		}
	}
	return false
}

func wrapRoute(router *mux.Router, f func(r *http.Request, vars map[string]string) string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !prod || isShort(r.Host) {
			if r.URL.Query().Get("debug") == "true" {
				w.WriteHeader(200)
				resp := f(r, mux.Vars(r))
				fmt.Println(r.URL, "=>", resp)
				w.Write([]byte("Going to redirect to " + resp))
				return
			}
			resp := f(r, mux.Vars(r))
			fmt.Println(r.URL, "=>", resp)
			http.Redirect(w, r, resp, http.StatusFound)
		} else {
			w.WriteHeader(400)
			w.Write([]byte("Not configured as short url"))
		}
	}
}

func shortRoute(path string, r *mux.Router, f func(r *http.Request, vars map[string]string) string) {
	shortRoutes[path] = f

	r.HandleFunc(path, wrapRoute(r, f))
}

func main() {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(wrapRoute(r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + r.URL.EscapedPath()
	}))

	// handle / route
	shortRoute("/", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain
	})

	// We redirect bots onto itself to allow for better redirects
	shortRoute("/bot/{id}", r, func(r *http.Request, vars map[string]string) string {
		return "https://" + r.Host + "/" + vars["id"]
	})

	shortRoute("/bot/{id}/{path}", r, func(r *http.Request, vars map[string]string) string {
		return "https://" + r.Host + "/" + vars["id"] + "/" + vars["path"]
	})

	shortRoute("/bots/{id}", r, func(r *http.Request, vars map[string]string) string {
		return "https://" + r.Host + "/" + vars["id"]
	})

	shortRoute("/bots/{id}/{path}", r, func(r *http.Request, vars map[string]string) string {
		return "https://" + r.Host + "/" + vars["id"] + "/" + vars["path"]
	})

	// Bot redirects
	shortRoute("/{id}", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/bots/" + vars["id"]
	})

	// Invite
	shortRoute("/{id}/i", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/bots/" + vars["id"] + "/invite"
	})

	shortRoute("/{id}/inv", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/bots/" + vars["id"] + "/invite"
	})

	shortRoute("/{id}/invite", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/bots/" + vars["id"] + "/invite"
	})

	// Vote
	shortRoute("/{id}/v", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/bots/" + vars["id"] + "/vote"
	})

	shortRoute("/{id}/vote", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/bots/" + vars["id"] + "/vote"
	})

	// Packs
	shortRoute("/p/{id}", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/packs/" + vars["id"]
	})

	shortRoute("/{id}/p", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/packs/" + vars["id"]
	})

	shortRoute("/{id}/packs", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/packs/" + vars["id"]
	})

	// User routes
	shortRoute("/u/{id}", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/users/" + vars["id"]
	})

	shortRoute("/user/{id}", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/users/" + vars["id"]
	})

	shortRoute("/users/{id}", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/users/" + vars["id"]
	})

	shortRoute("/profile/{id}", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/users/" + vars["id"]
	})

	shortRoute("/profiles/{id}", r, func(r *http.Request, vars map[string]string) string {
		return mainDomain + "/users/" + vars["id"]
	})

	// Short API

	r.HandleFunc("/api/redirects.json", func(w http.ResponseWriter, r *http.Request) {
		var apiResp = make(map[string]string)

		for k, v := range shortRoutes {
			apiResp[k] = v(r, map[string]string{
				"id":   "%ID%",
				"path": "%PATH%",
			})
		}

		bytes, err := json.Marshal(apiResp)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error marshalling JSON"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)

		fmt.Println(r.URL, "=>", http.StatusOK)
	})

	http.ListenAndServe(":1010", r)
}
