package island

import (
	"fmt"
	//"log"
)

/* ================= BASE STRUCT =================== */
type Place struct {
	Name    string
	ID      int
	Descrip string
	/// UP
	Location *PlaceHolder
	/// DOWN
	SubSections *PlaceSet
	Occupants   *PersonSet
	Contents    *ItemSet
	Features    *ItemSet
	Events      []*Event
	// SPECIFIC
	Exits *PlaceSet
}

func NewPlace(name string, g *Game) *Place {
	p := &Place{
		Name:        name,
		Location:    &PlaceHolder{},
		SubSections: NewPlaceSet(),
		Occupants:   NewPersonSet(),
		Contents:    NewItemSet(),
		Exits:       NewPlaceSet(),
		Features:    NewItemSet(),
		Events:      make([]*Event, 0),
	}
	p.Register(g)
	return p
}

func (p *Place) TString() string {
	return p.Name
}

func (p1 *Place) Connect(p2 *Place) {
	p1.Exits.Add(p2)
	p2.Exits.Add(p1)
}

func (p1 *Place) Connect1W(p2 *Place) {
	p1.Exits.Add(p2)
}

func (p *Place) SpawnPlace(name string, g *Game) *Place {
	p2 := NewPlace(name, g)
	p2.Location.Set(p)
	p.SubSections.Add(p2)
	p.Connect(p2)
	return p2
}

func (p *Place) SpawnPlace1W(name string, g *Game) *Place {
	p2 := NewPlace(name, g)
	p2.Location.Set(p)
	p.SubSections.Add(p2)
	p.Connect1W(p2)
	return p2
}

func (p *Place) SpawnPlace0W(name string, g *Game) *Place {
	p2 := NewPlace(name, g)
	p2.Location.Set(p)
	p.SubSections.Add(p2)
	return p2
}

func (p *Place) SpawnPerson(name string, g *Game) *Person {
	p2 := NewPerson(name, g)
	p2.MoveTo(p, g)
	return p2
}

func (p *Place) SpawnItem(name string, g *Game) *Item {
	i := NewItem(name, g)
	p.Contents.Add(i)
	i.Location.Set(p)
	return i
}

func (p *Place) SpawnFeature(name string, g *Game) *Item {
	f := NewItem(name, g)
	f.Big = true
	p.Features.Add(f)
	f.Location.Set(p)
	return f
}

func (p *Place) SpawnEvent() *Event {
	e := NewEvent()
	p.Events = append(p.Events, e)
	return e
}

// ==================================================
type Person struct {
	Name string
	ID   int
	/// UP
	Location *PlaceHolder
	/// DOWN
	Contents *ItemSet
	Events   []*Event
	// SPECIFIC
	Dialogue string
}

func NewPerson(name string, g *Game) *Person {
	p := &Person{
		Name:     name,
		Contents: NewItemSet(),
		Location: &PlaceHolder{},
		Events:   make([]*Event, 0),
	}
	p.Register(g)
	return p
}

func (p *Person) TString() string {
	return p.Name
}

func (p *Person) SpawnItem(name string, g *Game) *Item {
	i := NewItem(name, g)
	p.PickUp(i, g)
	return i
}

func (p *Person) SpawnEvent() *Event {
	e := NewEvent()
	p.Events = append(p.Events, e)
	return e
}

func (p *Person) MoveTo(pl *Place, g *Game) {
	l1 := p.Location.Get(g)
	if l1 != nil {
		l1.Occupants.Drop(p)
	}
	p.Location.Set(pl)
	pl.Occupants.Add(p)
}

func (p *Person) PickUp(i *Item, g *Game) {
	i.Free(g)
	p.Contents.Add(i)
	i.Holder.Set(p)
}

func (p *Person) Drop(i *Item, g *Game) {
	i.Free(g)
	loc := p.Location.Get(g)
	if loc != nil {
		i.Location.Set(loc)
		loc.Contents.Add(i)
	}
}

func (p *Person) Talk(t *Person, g *Game) string {
	return fmt.Sprintf("%s greets %s.", p.Name, t.Name)
}

// ==================================================
type Item struct {
	Name string
	ID   int
	/// UP
	Location *PlaceHolder
	Holder   *PersonHolder
	/// DOWN
	// SPECIFIC
	Big          bool
	Interactions map[int]*Interaction
}

func (i *Item) TString() string {
	return i.Name
}

func NewItem(name string, g *Game) *Item {
	i := &Item{
		Name:         name,
		Location:     &PlaceHolder{},
		Holder:       &PersonHolder{},
		Interactions: make(map[int]*Interaction),
	}
	i.Register(g)
	return i
}

func (i *Item) Free(g *Game) {
	loc := i.Location.Get(g)
	if loc != nil {
		loc.Contents.Drop(i)
		loc.Features.Drop(i)
		i.Location.Set(nil)
	}
	hold := i.Holder.Get(g)
	if hold != nil {
		hold.Contents.Drop(i)
		i.Holder.Set(nil)
	}
}

func (i *Item) UseOn(t *Item, g *Game) string {
	result, ok := t.Interactions[i.ID]
	if !ok {
		return "Nothing happens!"
	}
	return result.Run(g)

}

func (i *Item) AddIntr(i2 *Item, intr *Interaction) {
	i.Interactions[i2.ID] = intr
}
