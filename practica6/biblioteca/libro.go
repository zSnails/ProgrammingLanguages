package main

import "fmt"

type Libro struct {
	Codigo     uint32 `json:"codigo"`
	Titulo     string `json:"titulo"`
	Autor      string `json:"autor"`
	Disponible bool   `json:"disponible"`
	Prestado   bool   `json:"-"`
}

func (l Libro) String() string {
	return fmt.Sprintf("Libro(%d, %s, %s, estado=%s)", l.Codigo, l.Titulo, l.Autor, l.Localizacion())
}

// TODO: hacer esto más bonito
func (l Libro) Localizacion() string {
	if !l.Disponible && l.Prestado {
		return "en manos de un socio"
	}

	return "guardado"
}

func NewLibro(codigo uint32, titulo, autor string) *Libro {
	return &Libro{
		Codigo:     codigo,
		Titulo:     titulo,
		Autor:      autor,
		Disponible: true,
		Prestado:   false,
	}
}
