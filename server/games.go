package main

import (
	"mule/island"
	"net/http"
)

var games = SetupGames()

func SetupGames() map[string]func(http.ResponseWriter, *http.Request) {
	g := make(map[string]func(http.ResponseWriter, *http.Request))
	g["island"] = island.ServeHTTP
	return g
}
