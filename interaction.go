package roomgame

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
	p1.Exits.Drop(p2)
	p2.Exits.Drop(p1)
	return
}
