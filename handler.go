package roomgame

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

const REDIR = "/"

const TEMP_DIR = "templates/roomgame/"
const TEMP_EXT = ".html"
const MAIN_TEMPLATE = "main" + TEMP_EXT

var template_list = []string{TEMP_DIR + MAIN_TEMPLATE}
var templates = template.Must(template.ParseFiles(template_list...))

func ServeHTTP(w http.ResponseWriter, r *http.Request, filename string) {
	// =============== POST ==============
	var err error
	if r.Method == "POST" {
		err = POST_Process(r, filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		redir := strings.Join(strings.Split(r.URL.Path, "/")[:3], "/")
		http.Redirect(w, r, redir, http.StatusFound)
		return
	} // ==========  GET  ===========
	GET_Process(w, filename)
}

func GameFileFor(r *http.Request) (string, error) {
	return strings.Split(r.URL.Path, "/")[1], nil
}

func GET_Process(w http.ResponseWriter, gameFile string) {
	game, err := Load(gameFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = templates.ExecuteTemplate(w, MAIN_TEMPLATE, game.Page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func POST_Process(r *http.Request, gameFile string) error {
	game, err := Load(gameFile)
	if err != nil {
		Log(err)
		return err
	}
	action := r.FormValue("action")
	if game.Over && action != "NEWGAME" {
		return nil
	}
	target, err := strconv.ParseInt(r.FormValue("target"), 0, 0)
	if err != nil {
		Log(err)
		return err
	}
	t := int(target)
	f, ok := ActionMap[action]
	if !ok {
		err = errors.New(fmt.Sprintf("Action %s not a valid action", action))
		Log(err)
		return err
	}
	err = f(game, t)
	if err != nil {
		Log(err)
		return err
	}
	err = game.Save(gameFile)
	if err != nil {
		Log(err)
		return err
	}
	return nil
}

var ActionMap = NewActionMap()

func NewActionMap() map[string]func(*Game, int) error {
	a := make(map[string]func(*Game, int) error)
	a["move"] = userMove
	a["pickup"] = userPickup
	a["drop"] = userDrop
	a["talk"] = userTalk
	a["interact"] = userInteract
	a["NEWGAME"] = userNEWGAME
	return a
}
