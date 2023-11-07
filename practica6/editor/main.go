package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	g := NewGroup()
	ga := NewGroup()
	t := NewText("Esto es una prueba de un Text que tiene varios hijos")
	t2 := NewText("Prueba 1")

	s := NewSquare(5)
	t3 := NewText("Prueba 2")

	ga.AddChildren(t)
	ga.AddChildren(t2)
	t.AddChildren(t3)

	ga.AddChildren(s)
	// Tanto los objetos como los grupos pueden tener hijos
	s.AddChildren(t3)

	// Aquí hay un grupo externo que contiene un grupo interno
	g.AddChildren(&ga)

	// Y ya, para mostrarlo escogí usar un json
	data, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
}
