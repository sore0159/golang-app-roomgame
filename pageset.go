package island

import (
	"fmt"
	//	"sort"
)

type PageData struct {
	Place          string
	PDescr         string
	User           string
	PCName         string
	Time           int
	Chatlog        []string
	History        []string
	Features       []string
	FeatureButtons []*Button
	ExitButtons    []*Button
	PeopleButtons  []*Button
	LocItemButtons []*Button
	PCItemButtons  []*Button
	NewGameButton  *Button
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
		PCName:         pc.Name,
		Time:           g.Time,
		Place:          pc_loc.Name,
		PDescr:         pc_loc.Descrip,
		Chatlog:        g.Chatlog,
		History:        make([]string, 0),
		Features:       make([]string, 0),
		ExitButtons:    make([]*Button, 0),
		PeopleButtons:  make([]*Button, 0),
		LocItemButtons: make([]*Button, 0),
		PCItemButtons:  make([]*Button, 0),
		FeatureButtons: make([]*Button, 0),
		NewGameButton: &Button{
			Action: "NEWGAME",
			Target: "0",
			Text:   "DELETE GAME",
			Hover:  "Start Completely Over",
		},
	}
	// HISTORY
	l := len(g.History)
	if l > 3 {
		l -= 3
	} else {
		l = 0
	}
	for _, event := range g.History[l:] {
		p.History = append(p.History, event)
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
	var pcItemFlag string
	for thing, _ := range pc.Contents.Get(g) {
		pcItemFlag = thing.Name
		p.PCItemButtons = append(p.PCItemButtons, &Button{
			Action: "drop",
			Target: fmt.Sprintf("%d", thing.ID),
			Text:   thing.Name,
			Hover:  "Drop " + thing.Name,
		})
	}
	//=============
	// FEATURES (DOORS)
	if pcItemFlag != "" {
		for thing, _ := range pc_loc.Features.Get(g) {
			p.FeatureButtons = append(p.FeatureButtons, &Button{
				Action: "interact",
				Target: fmt.Sprintf("%d", thing.ID),
				Text:   thing.Name,
				Hover:  fmt.Sprintf("Use %s on %s", pcItemFlag, thing.Name),
			})
		}
	} else {
		for thing, _ := range pc_loc.Features.Get(g) {
			p.Features = append(p.Features, thing.Name)
		}
	}

	//=============

	g.Page = p
}
