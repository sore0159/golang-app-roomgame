package main

import (
	"errors"
	"fmt"
	"mule/island"
	"net/http"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("ping")
	http.HandleFunc("/", parsePath)
	http.ListenAndServe(":8080", nil)
	fmt.Println("ping")
}

var games = SetupGames()

func SetupGames() map[string]func(http.ResponseWriter, *http.Request) {
	g := make(map[string]func(http.ResponseWriter, *http.Request))
	g["island"] = island.ServeHTTP
	return g
}

func parsePath(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathParts := strings.Split(path, "/")
	l := len(pathParts)
	if pathParts[l-1] == "" {
		l--
		pathParts = pathParts[:l]
	}
	var user string
	switch {
	case l == 1:
		serveMain(w, r)
		return
	case l > 1:
		user = pathParts[1]
		if !isAuth(r, user) {
			err := errors.New(fmt.Sprintf("%s is not an authorized user", user))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if l == 2 {
			serveUserIndex(w, r, user)
			return
		}
		game := pathParts[2]
		//		serveTempGame(w, r, game, user)
		if f, ok := games[game]; ok {
			f(w, r)
		} else {
			err := errors.New(fmt.Sprintf("Can't find game %s", game))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func isAuth(r *http.Request, name string) bool {
	for _, rn := range name {
		if !unicode.In(rn, unicode.L, unicode.N) {
			return false
		}
	}
	return true
}

func serveMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MAIN INDEX")
}

func serveUserIndex(w http.ResponseWriter, r *http.Request, user string) {
	fmt.Fprintf(w, "INDEX FOR USER %s", user)
}

func serveTempGame(w http.ResponseWriter, r *http.Request, game, user string) {
	if game == "err" {
		err := errors.New(fmt.Sprintf("Can't find game %s", game))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "TEMP GAME %s FOR USER %s", game, user)
}
