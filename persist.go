package roomgame

import (
	"encoding/gob"
	"os"
)

func (g *Game) Save() error {
	dataFile, err := os.Create(g.FileName)
	if err != nil {
		Log(err)
		return err
	}
	defer dataFile.Close()
	dataEncoder := gob.NewEncoder(dataFile)
	err = dataEncoder.Encode(g)
	if err != nil {
		Log(err)
		return err
	}
	return nil
}

func Load(fileName string) (*Game, error) {
	var g *Game
	dataFile, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			g = New(fileName)
			err = g.Save()
			if err != nil {
				Log(err)
				return nil, err
			}
			return g, nil
		}
		Log(err)
		return nil, err
	}
	g = BlankGame()
	defer dataFile.Close()
	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(g)
	if err != nil {
		Log(err)
		return nil, err
	}
	g.FileName = fileName
	return g, nil
}

func New(fileName string) *Game {
	g := BlankGame()
	g.SetupGame1(fileName)
	g.FileName = fileName
	return g
}
