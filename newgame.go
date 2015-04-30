package island

func (g *Game) SetupGame1(userName string) {
	g.User = userName
	i := NewPlace("Small Island", g)
	pc := i.SpawnPerson("You", g)
	g.PC.Set(pc)
	p1 := i.SpawnPlace1W("Room One", g)
	p2 := i.SpawnPlace1W("Room Two", g)
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
