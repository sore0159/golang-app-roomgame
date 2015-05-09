package island

import (
	"errors"
	"fmt"
	"strings"
)

type Game struct {
	User     string
	FileName string
	Time     int
	Chatlog  []string
	History  []string
	PC       *PersonHolder
	Page     *PageData
	Reg      *Registry
}

func BlankGame() *Game {
	return &Game{
		PC:      &PersonHolder{},
		Reg:     NewRegistry(),
		Chatlog: make([]string, 0),
		History: make([]string, 0),
	}
}

func userMove(g *Game, target int) error {
	pc := g.PC.Get(g)
	loc := pc.Location.Get(g)
	tar := g.GetPlace(target)
	if tar == nil || !loc.Exits.Has(tar) {
		return errors.New("Invalid move destination!")
	}
	g.tic()
	g.Witness(fmt.Sprintf("%s moves from %s to %s.", pc.Name, loc.Name, tar.Name))
	pc.Location.Set(tar)
	g.PageSet()
	return nil
}

func userPickup(g *Game, target int) error {
	pc := g.PC.Get(g)
	loc := pc.Location.Get(g)
	tar := g.GetItem(target)
	if tar == nil || !loc.Contents.Has(tar) || tar.Big {
		return errors.New("Invalid pickup target!")
	}
	drops := make([]string, 0)
	if conSet := pc.Contents.Get(g); conSet != nil {
		for i, _ := range conSet {
			pc.Drop(i, g)
			drops = append(drops, i.Name)
		}
	}
	g.tic()
	if len(drops) > 0 {
		g.Witness(fmt.Sprintf("%s drops %s and picks up %s.", pc.Name, strings.Join(drops, ", "), tar.Name))
	} else {
		g.Witness(fmt.Sprintf("%s picks up %s.", pc.Name, tar.Name))
	}
	pc.PickUp(tar, g)
	g.PageSet()
	return nil
}

func userDrop(g *Game, target int) error {
	pc := g.PC.Get(g)
	//loc := pc.Location.Get(g)
	tar := g.GetItem(target)
	if tar == nil || !pc.Contents.Has(tar) {
		return errors.New("Invalid drop target!")
	}
	g.tic()
	g.Witness(fmt.Sprintf("%s drops %s.", pc.Name, tar.Name))
	pc.Drop(tar, g)
	g.PageSet()
	return nil
}

func userTalk(g *Game, target int) error {
	pc := g.PC.Get(g)
	loc := pc.Location.Get(g)
	tar := g.GetPerson(target)
	if tar == nil || !loc.Occupants.Has(tar) {
		return errors.New("Invalid talk target!")
	}
	g.tic()
	g.Witness(tar.Talk(pc, g))
	g.PageSet()
	return nil
}

func userInteract(g *Game, target int) error {
	pc := g.PC.Get(g)
	loc := pc.Location.Get(g)
	tar := g.GetItem(target)
	items := pc.Contents.Get(g)
	var item *Item
	for i, _ := range items {
		item = i
	}
	if tar == nil || !loc.Features.Has(tar) || item == nil {
		return errors.New("Invalid interaction!")
	}
	g.tic()
	g.Witness(fmt.Sprintf("%s uses %s on %s", pc.Name, item.Name, tar.Name))
	g.Witness(item.UseOn(tar, g))
	g.PageSet()
	return nil
}

func (g *Game) Witness(event string) {
	g.Chatlog = append(g.Chatlog, event)
}

func (g *Game) tic() {
	for _, event := range g.Chatlog {
		g.History = append(g.History, fmt.Sprintf("(%d)%s", g.Time, event))
	}
	g.Time++
	g.Chatlog = make([]string, 0)
}

func userNEWGAME(g *Game, target int) error {
	*g = *New(g.FileName)
	return nil
}
