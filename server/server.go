package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var mainMux = mux.NewRouter()

func main() {
	SetupUsers()
	//MUX SETUP
	r := mainMux
	r.StrictSlash(true)
	r.HandleFunc("/", MainPage).Name("main")
	r.HandleFunc("/login", LoginPage).Name("login")
	r.HandleFunc("/login/{rest:.*}", RedirPage("login"))
	r.HandleFunc("/logout", LogoutPage).Name("logout")
	r.HandleFunc("/logout/{rest:.*}", RedirPage("logout"))
	r.HandleFunc("/{user:[a-zA-Z0-9]+}", authChecked(UserPage)).Name("user")
	r.HandleFunc("/{user:[a-zA-Z0-9]+}/{game:[a-zA-Z]+}", authChecked(GamePage)).Name("game")
	r.HandleFunc("/{user:[a-zA-Z0-9]+}/{game:[a-zA-Z]+}/{rest:.*}", RedirPage("game"))
	// OKAY GO!
	http.Handle("/", r)
	log.Println("STARTING SERVER")
	http.ListenAndServe(":8080", nil)
}

func reverse(name string, args ...string) string {
	url, err := mainMux.Get(name).URL(args...)
	if err != nil {
		log.Println(err)
		return "/"
	} else {
		return url.Path
	}
}

func RedirPage(page string) func(http.ResponseWriter, *http.Request) {
	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		v := make([]string, 0)
		for key, val := range vars {
			if key == "rest" {
				continue
			}
			v = append(v, key, val)
		}
		http.Redirect(w, r, reverse(page, v...), http.StatusFound)
		return
	}
	return f
}
