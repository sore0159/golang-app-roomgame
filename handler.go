package island

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const REDIR = "/"

const TEMP_DIR = "templates/"
const TEMP_EXT = ".html"
const ISLAND_TEMPLATE = "island" + TEMP_EXT

var template_list = []string{TEMP_DIR + ISLAND_TEMPLATE}
var templates = template.Must(template.ParseFiles(template_list...))

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ======== LOOK AT r TO FIND WHICH GAME ========
	gameFile, err := GameFileFor(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// =============== POST ==============
	if r.Method == "POST" {
		//log.Println("POST")
		err = POST_Process(r, gameFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		redir := strings.Join(strings.Split(r.URL.Path, "/")[:3], "/")
		http.Redirect(w, r, redir, http.StatusFound)
		return
	} // ==========  GET  ===========
	GET_Process(w, gameFile)
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
	err = templates.ExecuteTemplate(w, ISLAND_TEMPLATE, game.Page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func POST_Process(r *http.Request, gameFile string) error {
	game, err := Load(gameFile)
	if err != nil {
		log.Println(err)
		return err
	}
	action := r.FormValue("action")
	target, err := strconv.ParseInt(r.FormValue("target"), 0, 0)
	if err != nil {
		log.Println(err)
		return err
	}
	t := int(target)
	f, ok := ActionMap[action]
	if !ok {
		err = errors.New(fmt.Sprintf("Action %s not a valid action", action))
		log.Println(err)
		return err
	}
	err = f(game, t)
	if err != nil {
		log.Println(err)
		return err
	}
	err = game.Save(gameFile)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

var ActionMap = NewActionMap()

func NewActionMap() map[string]func(*Game, int) error {
	a := make(map[string]func(*Game, int) error)
	a["move"] = userMove
	return a
}
