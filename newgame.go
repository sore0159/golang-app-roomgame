package roomgame

import (
	"mule/randomlists"
	"strings"
)

func (g *Game) SetupGame2(fileName string) {
	user := userName(fileName)
	g.User = user
	places := randomlists.Rooms(4)
	people := randomlists.People(3)
	i := NewPlace(places[0], g)

	pc := i.SpawnPerson(strings.Title(strings.ToLower(user)), g)
	g.PC.Set(pc)
	p1 := i.SpawnPlace1W(places[1], g)
	p2 := i.SpawnPlace1W(places[2], g)
	p3 := i.SpawnPlace0W(places[3], g)
	p3.Descrip = "This room is pretty awesome."
	gk := i.SpawnItem("Green Key", g)
	g.LockedDoor("Green", p2, p3, gk)
	p1.Connect(p2)
	p1.SpawnItem("Fork", g)
	p2.SpawnItem("Shovel", g)
	p2.SpawnItem("Book", g)
	p2.SpawnPerson(people[0], g)
	p2.SpawnPerson(people[1], g)
	pc.SpawnItem("Hat", g)
	tim := p3.SpawnPerson(people[2], g)
	tim.SpawnItem("Cloak", g)
	g.Objectives.Add(tim)
	g.PageSet()
}

func userName(fileName string) string {
	user := strings.Split(fileName, "/")[3]
	if len(user) > 5 && user[:6] == "guest-" {
		return "Guest"
	}
	return user
}
