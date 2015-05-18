package roomgame

import (
	"errors"
	"fmt"
	"html/template"
	"mule/apps/auth"
	"net/http"
	"strconv"
)

const REDIR = "/"

const TEMP_DIR = "templates/" + APPNAME + "/"
const TEMP_EXT = ".html"
const MAIN_TEMPLATE = "main" + TEMP_EXT

var template_list = []string{TEMP_DIR + MAIN_TEMPLATE}
var templates = template.Must(template.ParseFiles(template_list...))

func RoomPage(d *auth.Data, g *Game) {
	// =============== POST ==============
	if d.R.Method == "POST" {
		err := POST_Process(d.R, g)
		if err != nil {
			http.Error(d.W, err.Error(), http.StatusInternalServerError)
			return
		}
		d.GoGame()
		return
	} // ==========  GET  ===========
	GET_Process(d, g)
}

func GET_Process(d *auth.Data, g *Game) {
	g.Page.UName = d.Name
	g.Page.HomeURL = d.HomeURL
	g.Page.GameURL = d.GameURL
	err := templates.ExecuteTemplate(d.W, MAIN_TEMPLATE, g.Page)
	if err != nil {
		http.Error(d.W, err.Error(), http.StatusInternalServerError)
	}
}

func POST_Process(r *http.Request, game *Game) error {
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
	err = game.Save()
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
