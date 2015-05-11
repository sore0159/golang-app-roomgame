package roomgame

import (
	"strings"
)

func (g *Game) SetupGame2(fileName string) {
	user := userName(fileName)
	g.User = user
	i := NewPlace("Tiny Game", g)
	pc := i.SpawnPerson(strings.Title(strings.ToLower(user)), g)
	g.PC.Set(pc)
	p1 := i.SpawnPlace1W("Room One", g)
	p2 := i.SpawnPlace1W("Room Two", g)
	p3 := i.SpawnPlace0W("Room Three", g)
	p3.Descrip = "This room is pretty awesome."
	gk := i.SpawnItem("Green Key", g)
	g.LockedDoor("Green", p2, p3, gk)
	p1.Connect(p2)
	p1.SpawnItem("Fork", g)
	p2.SpawnItem("Shovel", g)
	p2.SpawnItem("Book", g)
	p2.SpawnPerson("Anon", g)
	p2.SpawnPerson("Mous", g)
	pc.SpawnItem("Hat", g)
	tim := p1.SpawnPerson("Tim", g)
	tim.SpawnItem("Cloak", g)
	g.PageSet()
}

func userName(fileName string) string {
	user := strings.Split(fileName, "/")[3]
	if len(user) > 5 && user[:6] == "guest-" {
		return "Guest"
	}
	return user
}
