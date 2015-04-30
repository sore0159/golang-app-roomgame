package island

import (
	"fmt"
	//	"sort"
)

type PageData struct {
	Place          string
	User           string
	Time           int
	ExitButtons    []*Button
	PeopleButtons  []*Button
	LocItemButtons []*Button
	PCItemButtons  []*Button
}

type Button struct {
	Action string
	Target string
	Text   string
	Hover  string
}

func (g *Game) PageSet() {
	pc := g.PC.Get(g)
	pc_loc := pc.Location.Get(g)
	p := &PageData{
		User:           g.User,
		Time:           g.Time,
		Place:          pc_loc.Name,
		ExitButtons:    make([]*Button, 0),
		PeopleButtons:  make([]*Button, 0),
		LocItemButtons: make([]*Button, 0),
		PCItemButtons:  make([]*Button, 0),
	}
	// EXIT BUTTONS
	for loc, _ := range pc_loc.Exits.Get(g) {
		p.ExitButtons = append(p.ExitButtons, &Button{
			Action: "move",
			Target: fmt.Sprintf("%d", loc.ID),
			Text:   loc.Name,
			Hover:  "Move to " + loc.Name,
		})
	}
	//=============
	// PEEP BUTTONS
	for peep, _ := range pc_loc.Occupants.Get(g) {
		if peep == pc {
			continue
		}
		p.PeopleButtons = append(p.PeopleButtons, &Button{
			Action: "talk",
			Target: fmt.Sprintf("%d", peep.ID),
			Text:   peep.Name,
			Hover:  "Talk to " + peep.Name,
		})
	}
	//=============
	// LOC ITEM BUTTONS
	for thing, _ := range pc_loc.Contents.Get(g) {
		p.LocItemButtons = append(p.LocItemButtons, &Button{
			Action: "pickup",
			Target: fmt.Sprintf("%d", thing.ID),
			Text:   thing.Name,
			Hover:  "Pick Up " + thing.Name,
		})
	}
	//=============
	// PC ITEM BUTTONS
	for thing, _ := range pc.Contents.Get(g) {
		p.PCItemButtons = append(p.PCItemButtons, &Button{
			Action: "drop",
			Target: fmt.Sprintf("%d", thing.ID),
			Text:   thing.Name,
			Hover:  "Drop " + thing.Name,
		})
	}
	//=============

	g.Page = p
}
