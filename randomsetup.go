package roomgame

import (
	"fmt"
	"math/rand"
	"mule/randomlists"
	"strings"
)

var Seed = randomlists.Seed

type Section struct {
	unconnected *PlaceSet
	all         *PlaceSet
	endpoints   *PlaceSet
}

type Level struct {
	part1       []*Section
	part2       []*Section
	key         *Item
	food        *Item
	unconnected *PlaceSet
	endpoints1  *PlaceSet
	endpoints2  *PlaceSet
	color       string
	all         *PlaceSet
}

func NewSection(g *Game) *Section {
	Seed()
	s := &Section{
		endpoints: NewPlaceSet(),
	}
	numPlaces := 4 + rand.Intn(3)
	m := NewPlaceSet()
	m2 := NewPlaceSet()
	var lastPlace *Place
	var newPlace *Place
	for i := 0; i < numPlaces; i++ {
		newPlace = NewPlace("temp", g)
		m2.Add(newPlace)
		if lastPlace != nil {
			newPlace.Connect(lastPlace)
			m.Add(newPlace)
		} else {
			s.endpoints.Add(newPlace)
		}
		lastPlace = newPlace
	}
	s.endpoints.Add(newPlace)
	m.Drop(newPlace)
	s.unconnected = m
	s.all = m2
	return s
}

func (g *Game) randSet(s *PlaceSet) *Place {
	Seed()
	m := s.Get(g)
	l := len(m)
	if l < 1 {
		Log(s)
	}
	p := rand.Intn(l)
	i := 0
	for v, _ := range m {
		if i == p {
			return v
		}
		i++
	}
	return nil
}

func (g *Game) ConnectSections(s1, s2 *Section) {
	p1 := g.randSet(s1.unconnected)
	p2 := g.randSet(s2.unconnected)
	b := g.randSet(p1.Exits)
	p1.UnConnect(b)
	p2.Connect(b)
	p2.Connect(p1)
	s2.unconnected.Drop(p2)
}
func (s *Section) leng(g *Game) int {
	return len(s.all.Get(g))
}
func (l *Level) leng(g *Game) int {
	var t int
	for _, s := range l.part1 {
		t += s.leng(g)
	}
	for _, s := range l.part2 {
		t += s.leng(g)
	}
	return t
}

func NewLevel(color string, g *Game) *Level {
	k := NewItem(color+" key", g)
	f := NewItem("tempfood", g)
	lvl := &Level{
		key:         k,
		food:        f,
		color:       color,
		part1:       make([]*Section, 0),
		unconnected: NewPlaceSet(),
		endpoints1:  NewPlaceSet(),
		endpoints2:  NewPlaceSet(),
		all:         NewPlaceSet(),
	}
	Seed()
	l1 := 2 + rand.Intn(1)
	l2 := rand.Intn(1)
	if l2 == 1 {
		l1 -= 1
	}
	for i := 0; i < l1; i++ {
		s := NewSection(g)
		for p, _ := range s.all.Get(g) {
			lvl.all.Add(p)
		}
		lvl.part1 = append(lvl.part1, s)
		if i != 0 {
			g.ConnectSections(lvl.part1[i-1], s)
		}
	}
	for _, s := range lvl.part1 {
		for p, _ := range s.unconnected.Get(g) {
			lvl.unconnected.Add(p)
		}
		for p, _ := range s.endpoints.Get(g) {
			lvl.endpoints1.Add(p)
		}
	}
	if l2 == 1 {
		p2 := make([]*Section, 1)
		s := NewSection(g)
		for p, _ := range s.all.Get(g) {
			lvl.all.Add(p)
		}
		p2[0] = s
		lvl.part2 = p2
		for p, _ := range s.unconnected.Get(g) {
			lvl.unconnected.Add(p)
		}
		for p, _ := range s.endpoints.Get(g) {
			lvl.endpoints2.Add(p)
		}
	}
	return lvl
}

func (g *Game) ConnectLevels(lower, higher *Level) {
	all := NewPlaceSet()
	for x, _ := range higher.endpoints1.Get(g) {
		all.Add(x)
	}
	if higher.part2 != nil {
		for x, _ := range higher.endpoints2.Get(g) {
			all.Add(x)
		}
	}
	numEnd := len(all.Get(g))
	for i := 0; i < numEnd/2; i++ {
		if i == 0 && higher.part2 != nil {
			p1 := g.randSet(higher.endpoints1)
			all.Drop(p1)
			p2 := g.randSet(higher.endpoints2)
			all.Drop(p2)
			l1 := g.randSet(lower.unconnected)
			lower.unconnected.Drop(l1)
			l2 := g.randSet(lower.unconnected)
			lower.unconnected.Drop(l2)
			g.LockedDoor(higher.color, l1, p1, higher.key)
			g.LockedDoor(higher.color, l2, p2, higher.key)
			continue
		}
		p := g.randSet(all)
		all.Drop(p)
		l := g.randSet(lower.unconnected)
		lower.unconnected.Drop(l)
		g.LockedDoor(higher.color, l, p, higher.key)
	}
	var kPlace *Place
	if len(lower.unconnected.Get(g)) < 1 {
		kPlace = g.randSet(lower.endpoints1)
	} else {
		kPlace = g.randSet(lower.unconnected)
	}
	kPlace.Contents.Add(higher.key)
	higher.key.Location.Set(kPlace)
	// Food?
	// Items of Power?
}

func (g *Game) SetupGame1(fileName string) {
	user := userName(fileName)
	g.User = user
	numLevels := 3 + rand.Intn(3)
	colors := randomlists.Colors(numLevels)
	peopleN := randomlists.People(numLevels + 1)
	levels := make([]*Level, numLevels)
	for i, _ := range levels {
		levels[i] = NewLevel(colors[i], g)
	}
	//
	placeNames := randomlists.Rooms(g.Reg.NumPlaces)
	var i int
	for _, place := range g.Reg.Places {
		place.Name = placeNames[i]
		i++
	}
	startSpot := NewPlace("Start Spot", g)
	greeter := startSpot.SpawnPerson(peopleN[numLevels], g)
	str1 := fmt.Sprintf("%d", numLevels+1)
	str2 := fmt.Sprintf("%d", numLevels)
	greeter.Speech = "%s welcomes %s to RoomGame Go!  Find all " + str1 + " people (" + str2 + " now) and talk with them to win!"
	g.Objectives.Add(greeter)
	pc := startSpot.SpawnPerson(strings.Title(strings.ToLower(user)), g)
	g.PC.Set(pc)
	x := g.randSet(levels[0].unconnected)
	startSpot.Connect(x)
	levels[0].unconnected.Drop(x)
	if levels[0].part2 != nil {
		p2 := g.randSet(levels[0].endpoints2)
		p1 := g.randSet(levels[0].endpoints1)
		levels[0].endpoints2.Drop(p2)
		levels[0].endpoints1.Drop(p1)
		p2.Connect(p1)
	}
	//
	keys := make([]*Item, 0)
	for i, lvl := range levels {
		if i != 0 {
			g.ConnectLevels(levels[i-1], lvl)
			g.lockup(lvl, keys)
		}
		keys = append(keys, lvl.key)
		p := g.randSet(lvl.all)
		per := p.SpawnPerson(peopleN[i], g)
		g.Objectives.Add(per)
	}
	g.PageSet()
}

func (g *Game) lockup(lvl *Level, keys []*Item) {
	for _, key := range keys {
		j := 1 + rand.Intn(2)
		for i := 0; i < j; i++ {
			var l int
			var p1 *Place
			for l == 0 {
				p1 = g.randSet(lvl.all)
				l = len(p1.Exits.Get(g))
			}
			p2 := g.randSet(p1.Exits)
			g.LockedDoor(lvl.color, p1, p2, key)
		}
	}
}
