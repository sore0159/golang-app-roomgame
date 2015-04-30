package island

import (
//"log"
)

/* ================= BASE STRUCT =================== */
type Place struct {
	Name string
	ID   int
	/// UP
	Location *PlaceHolder
	/// DOWN
	SubSections *PlaceSet
	Occupants   *PersonSet
	Contents    *ItemSet
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
	p.Connect(p2)
	return p2
}

func (p *Place) SpawnPlace1W(name string, g *Game) *Place {
	p2 := NewPlace(name, g)
	p.Connect1W(p2)
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

// ==================================================
type Person struct {
	Name string
	ID   int
	/// UP
	Location *PlaceHolder
	/// DOWN
	Contents *ItemSet
	// SPECIFIC
}

func NewPerson(name string, g *Game) *Person {
	p := &Person{
		Name:     name,
		Contents: NewItemSet(),
		Location: &PlaceHolder{},
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

func (p *Person) MoveTo(pl *Place, g *Game) {
	l1 := p.Location.Get(g)
	if l1 != nil {
		l1.Occupants.Drop(p)
	}
	p.Location.Set(pl)
	pl.Occupants.Add(p)
}

func (p *Person) PickUp(i *Item, g *Game) {
	hold := i.Holder.Get(g)
	if hold != nil {
		hold.Contents.Drop(i)
	}
	loc := i.Location.Get(g)
	if loc != nil {
		loc.Contents.Drop(i)
		i.Location.Set(nil)
	}
	i.Holder.Set(p)
	p.Contents.Add(i)
}

func (p *Person) Drop(i *Item, g *Game) {
	loc := p.Location.Get(g)
	p.Contents.Drop(i)
	i.Holder.Set(nil)
	i.Location.Set(loc)
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
}

func (i *Item) TString() string {
	return i.Name
}

func NewItem(name string, g *Game) *Item {
	i := &Item{
		Name:     name,
		Location: &PlaceHolder{},
		Holder:   &PersonHolder{},
	}
	i.Register(g)
	return i
}
