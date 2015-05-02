package main

import (
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var users map[string]bool

const MAX_USERS = 50

//  USER INFO
// data/users/NAME/
//			lastpost		Time/IP addr of last POST request
//			logs.txt		Logs (copy) for actions by this user
//			errlogs.txt		err logs
//			profile.txt		Profile info?
//			saves/
//					GAMENAME.gob		save data for games

func isAuth(r *http.Request, name string) bool {
	// check r to make sure user has access to area
	path := r.URL.Path
	_ = path
	return users[name]
}

func authChecked(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	f2 := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user := vars["user"]
		if isAuth(r, user) {
			f(w, r)
		} else {
			// IF LOGGED IN
			//		http.Redirect(w, r, reverse("user", "user", loggeduser), http.StatusFound)
			// ELSE
			// NOT LOGGED IN
			http.Redirect(w, r, reverse("login"), http.StatusFound)
		}
	}
	return f2
}

func SetupUsers() {
	userdir, err := os.Open("data/users")
	if err != nil {
		log.Fatal(err)
	}
	userlist, err := userdir.Readdirnames(MAX_USERS)
	if err != nil {
		log.Println(err)
	}
	l := len(userlist)

	if l < 1 {
		log.Fatal(errors.New("No users found!"))
	}
	log.Println("Updating Users:", userlist)
	m := make(map[string]bool, l)
	for _, u := range userlist {
		m[u] = true
	}
	users = m
}

func MakeUser(name string) error {
	userdir, err := os.Open("data/users")
	if err != nil {
		log.Fatal(err)
	}
	userlist, err := userdir.Readdirnames(MAX_USERS)
	if err != nil {
		log.Println(err)
	}
	l := len(userlist)
	if l == MAX_USERS {
		return errors.New("Too many users, cannot create new user!")
	}
	cmd := exec.Command("cp", "-r", "data/user_skel", "data/users/"+name)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
