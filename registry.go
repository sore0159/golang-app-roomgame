package island

import (
//"log"
//	"errors"
)

const ID_PR = 1
const ID_PL = 2
const ID_IT = 3
const ID_MAG = 10

type Registry struct {
	NumPlaces int
	Places    map[int]*Place
	NumPeople int
	People    map[int]*Person
	NumItems  int
	Items     map[int]*Item
}

func NewRegistry() *Registry {
	return &Registry{
		Places: make(map[int]*Place),
		People: make(map[int]*Person),
		Items:  make(map[int]*Item),
	}
}

func (p *Person) Register(g *Game) {
	g.Reg.NumPeople++
	p.ID = g.Reg.NumPeople*ID_MAG + ID_PR
	g.Reg.People[p.ID] = p
}

func (p *Place) Register(g *Game) {
	g.Reg.NumPlaces++
	p.ID = g.Reg.NumPlaces*ID_MAG + ID_PL
	g.Reg.Places[p.ID] = p
}

func (i *Item) Register(g *Game) {
	g.Reg.NumItems++
	i.ID = g.Reg.NumItems*ID_MAG + ID_IT
	g.Reg.Items[i.ID] = i
}

func (g *Game) GetPerson(id int) *Person {
	if id%ID_MAG != ID_PR {
		return nil
	}
	return g.Reg.People[id]
}

func (g *Game) GetItem(id int) *Item {
	if id%ID_MAG != ID_IT {
		return nil
	}
	return g.Reg.Items[id]
}

func (g *Game) GetPlace(id int) *Place {
	if id%ID_MAG != ID_PL {
		return nil
	}
	return g.Reg.Places[id]
}

func (p *Person) UnRegister(g *Game) {
	delete(g.Reg.People, p.ID)
}

func (p *Place) UnRegister(g *Game) {
	delete(g.Reg.Places, p.ID)
}

func (i *Item) UnRegister(g *Game) {
	delete(g.Reg.Items, i.ID)
}

//============ FAKE SETS ================
// ======= PERSON ========

type PersonSet struct {
	Data  map[int]bool
	cache map[*Person]bool
}

func NewPersonSet() *PersonSet {
	return &PersonSet{Data: make(map[int]bool)}
}

func (s *PersonSet) Add(p *Person) {
	s.Data[p.ID] = true
	if s.cache != nil {
		s.cache[p] = true
	}
}
func (s *PersonSet) Drop(p *Person) {
	delete(s.Data, p.ID)
	if s.cache != nil {
		delete(s.cache, p)
	}
}
func (s *PersonSet) Has(p *Person) bool {
	return s.Data[p.ID]
}
func (s *PersonSet) Get(g *Game) map[*Person]bool {
	if s.cache == nil {
		s.cache = make(map[*Person]bool)
		for key, _ := range s.Data {
			s.cache[g.GetPerson(key)] = true
		}
	}
	return s.cache
}

type PersonHolder struct {
	Data  int
	cache *Person
}

func (h *PersonHolder) Set(p *Person) {
	h.Data = p.ID
	h.cache = p
}

func (h *PersonHolder) Get(g *Game) *Person {
	if h.cache == nil {
		h.cache = g.GetPerson(h.Data)
	}
	return h.cache
}

// ======= PLACE ========

type PlaceSet struct {
	Data  map[int]bool
	cache map[*Place]bool
}

func NewPlaceSet() *PlaceSet {
	return &PlaceSet{Data: make(map[int]bool)}
}

func (s *PlaceSet) Add(p *Place) {
	s.Data[p.ID] = true
	if s.cache != nil {
		s.cache[p] = true
	}
}
func (s *PlaceSet) Drop(p *Place) {
	delete(s.Data, p.ID)
	if s.cache != nil {
		delete(s.cache, p)
	}
}
func (s *PlaceSet) Has(p *Place) bool {
	return s.Data[p.ID]
}
func (s *PlaceSet) Get(g *Game) map[*Place]bool {
	if s.cache == nil {
		s.cache = make(map[*Place]bool)
		for key, _ := range s.Data {
			s.cache[g.GetPlace(key)] = true
		}
	}
	return s.cache
}

type PlaceHolder struct {
	Data  int
	cache *Place
}

func (h *PlaceHolder) Set(p *Place) {
	h.Data = p.ID
	h.cache = p
}

func (h *PlaceHolder) Get(g *Game) *Place {
	if h.cache == nil {
		h.cache = g.GetPlace(h.Data)
	}
	return h.cache
}

// ======= ITEM ========

type ItemSet struct {
	Data  map[int]bool
	cache map[*Item]bool
}

func NewItemSet() *ItemSet {
	return &ItemSet{Data: make(map[int]bool)}
}

func (s *ItemSet) Add(i *Item) {
	s.Data[i.ID] = true
	if s.cache != nil {
		s.cache[i] = true
	}
}
func (s *ItemSet) Drop(i *Item) {
	delete(s.Data, i.ID)
	if s.cache != nil {
		delete(s.cache, i)
	}
}
func (s *ItemSet) Has(i *Item) bool {
	return s.Data[i.ID]
}
func (s *ItemSet) Get(g *Game) map[*Item]bool {
	if s.cache == nil {
		s.cache = make(map[*Item]bool)
		for key, _ := range s.Data {
			s.cache[g.GetItem(key)] = true
		}
	}
	return s.cache
}

type ItemHolder struct {
	Data  int
	cache *Item
}

func (h *ItemHolder) Set(i *Item) {
	h.Data = i.ID
	h.cache = i
}

func (h *ItemHolder) Get(g *Game) *Item {
	if h.cache == nil {
		h.cache = g.GetItem(h.Data)
	}
	return h.cache
}