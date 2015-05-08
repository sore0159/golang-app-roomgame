package island

import (
	"fmt"
)

// ==================================================
type Interaction struct {
	Trigger    *ItemHolder
	Descrip    string
	Connects2W []*PlaceHolder
	Destroys   *ItemSet
}

func NewInteraction() *Interaction {
	return &Interaction{
		Trigger:    &ItemHolder{},
		Connects2W: make([]*PlaceHolder, 0),
		Destroys:   NewItemSet(),
	}
}

func (i *Interaction) Run(g *Game) string {
	for item, _ := range i.Destroys.Get(g) {
		item.Free(g)
	}
	for key, placeH := range i.Connects2W {
		if key%2 == 0 {
			placeH.Get(g).Connect(i.Connects2W[key+1].Get(g))
		}
	}
	return i.Descrip
}

func (i *Interaction) Add2WConn(p1, p2 *Place) {
	i.Connects2W = append(i.Connects2W, p1.NewHolder(), p2.NewHolder())
}

func (g *Game) LockedDoor(doorColor string, p1, p2 *Place, key *Item) (d1, d2 *Item) {
	d1 = p1.SpawnFeature(fmt.Sprintf("locked %s door to %s", doorColor, p2.Name), g)
	d2 = p2.SpawnFeature(fmt.Sprintf("locked %s door to %s", doorColor, p1.Name), g)
	intr := NewInteraction()
	intr.Trigger.Set(key)
	intr.Descrip = fmt.Sprintf("The %s door unlocks!", doorColor)
	intr.Add2WConn(p1, p2)
	intr.Destroys.Add(d1)
	intr.Destroys.Add(d2)
	d1.AddIntr(key, intr)
	d2.AddIntr(key, intr)
	return
}

// ==========================
type Condition struct {
	D bool
	I *ItemSet
	P *PersonSet
	L *PlaceHolder
}

func NewCondition() *Condition {
	return &Condition{
		D: true,
		I: NewItemSet(),
		P: NewPersonSet(),
		L: &PlaceHolder{},
	}
}

// If Loc, checks all P+I are at loc
// If not, needs 1 P, checks all I held by P
func (c *Condition) IsMet(g *Game) bool {
	items := c.I.Get(g)
	peeps := c.P.Get(g)
	loc := c.L.Get(g)
	if loc != nil {
		for peep, _ := range peeps {
			pLoc := peep.Location.Get(g)
			if pLoc != loc {
				return !c.D
			}
		}
		for item, _ := range items {
			iLoc := item.Location.Get(g)
			if iLoc != loc {
				return !c.D
			}
		}
	} else {
		if len(peeps) != 1 {
			return !c.D
		}
		for item, _ := range items {
			iHold := item.Holder.Get(g)
			if !peeps[iHold] {
				return !c.D
			}
		}
	}
	return c.D
}

type Action struct {
	Descrip string
}

func NewAction() *Action {
	return &Action{
	// stuff
	}
}

func (a *Action) Run(g *Game) string {
	return a.Descrip
}

type Event struct {
	PreConditions  []*Condition
	PostConditions []*Condition
	// All must be satisfied
	Results []*Action
}

func NewEvent() *Event {
	return &Event{
		PreConditions:  make([]*Condition, 0),
		PostConditions: make([]*Condition, 0),
		Results:        make([]*Action, 0),
	}
}

func (e *Event) PreCheck(g *Game) bool {
	for _, c := range e.PreConditions {
		if !c.IsMet(g) {
			return false
		}
	}
	return true
}

func (e *Event) PostCheck(g *Game) bool {
	for _, c := range e.PostConditions {
		if !c.IsMet(g) {
			return false
		}
	}
	return true
}

func (e *Event) Run(g *Game) []string {
	res := make([]string, 0)
	for _, a := range e.Results {
		res = append(res, a.Run(g))
	}
	return res
}

func (e *Event) SpawnPreC() *Condition {
	c := NewCondition()
	e.PreConditions = append(e.PreConditions, c)
	return c
}

func (e *Event) SpawnPostC() *Condition {
	c := NewCondition()
	e.PostConditions = append(e.PostConditions, c)
	return c
}

func (e *Event) SpawnAct() *Action {
	a := NewAction()
	e.Results = append(e.Results, a)
	return a
}
