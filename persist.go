package island

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

const SV_DIR = "data/users/%s/saves/island"
const SV_EXT = ".gob"

func SaveFull(name string) string {
	return fmt.Sprintf(SV_DIR, name) + SV_EXT
}

func (g *Game) Save(fileName string) error {
	dataFile, err := os.Create(SaveFull(fileName))
	if err != nil {
		log.Println(err)
		return err
	}
	defer dataFile.Close()
	dataEncoder := gob.NewEncoder(dataFile)
	err = dataEncoder.Encode(g)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Load(fileName string) (*Game, error) {
	var g *Game
	dataFile, err := os.Open(SaveFull(fileName))
	if err != nil {
		if os.IsNotExist(err) {
			g = New(fileName)
			err = g.Save(fileName)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			return g, nil
		}
		log.Println(err)
		return nil, err
	}
	g = BlankGame()
	defer dataFile.Close()
	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(g)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return g, nil
}

func New(fileName string) *Game {
	g := BlankGame()
	g.SetupGame1(fileName)
	return g
}
