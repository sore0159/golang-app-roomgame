package island

import (
	//	"fmt"
	"errors"
)

type Game struct {
	User string
	Time int
	PC   *PersonHolder
	Page *PageData
	Reg  *Registry
}

func BlankGame() *Game {
	return &Game{
		PC:  &PersonHolder{},
		Reg: NewRegistry(),
	}
}

func userMove(g *Game, target int) error {
	pc := g.PC.Get(g)
	loc := pc.Location.Get(g)
	tar := g.GetPlace(target)
	if tar != nil && loc.Exits.Has(tar) {
		g.Time++
		pc.Location.Set(tar)
		g.PageSet()
		return nil
	}
	return errors.New("Invalid move destination!")
}
