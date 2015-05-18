package roomgame

import (
	"github.com/gorilla/mux"
	"mule/apps/auth"
	"net/http"
)

const APPNAME = "roomgame"

func SetupMux(r *mux.Router) {
	r.HandleFunc("/", gWrap(RoomPage))
	r.HandleFunc("/{rest:.*}", auth.RedirGame(APPNAME))
}

func gWrap(f func(*auth.Data, *Game)) func(http.ResponseWriter, *http.Request) {
	return auth.DataWrap(auth.GameURLWrap(APPNAME, GameWrap(f)))
}

func GameWrap(f func(*auth.Data, *Game)) func(*auth.Data) {
	g := func(d *auth.Data) {
		dirName, err := d.SaveDir(APPNAME)
		if err != nil {
			d.ErrLog(err)
			http.Error(d.W, err.Error(), http.StatusInternalServerError)
			return
		}
		fileName := dirName + "/" + APPNAME + ".gob"
		gameData, err := Load(fileName)
		if err != nil {
			d.ErrLog(err)
			http.Error(d.W, err.Error(), http.StatusInternalServerError)
			return
		}
		f(d, gameData)
	}
	return g
}
