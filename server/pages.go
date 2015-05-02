package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if user, ok := validLogin(r); ok {
			// DO LOGIN STUFF
			http.Redirect(w, r, reverse("user", "user", user), http.StatusFound)
			return
		} else {
			// ERROR FLASH MESSAGE
			http.Redirect(w, r, reverse("login"), http.StatusFound)
			return
		}
	}
	fmt.Fprintf(w, "LOGIN PAGE")
}

func validLogin(r *http.Request) (string, bool) {
	user := ""
	return user, true
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// DO LOGOUT
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	fmt.Fprintf(w, "INDEX FOR USER %s", user)
}

func GamePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	game := vars["game"]
	if game == "err" {
		err := errors.New(fmt.Sprintf("Can't find game %s", game))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "TEMP GAME %s FOR USER %s", game, user)
}
func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MAIN INDEX")
}
