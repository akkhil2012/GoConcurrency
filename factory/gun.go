package main

type Gun struct {
	name string
}

func (g *Gun) setName(name string) {
	g.name = name
}
