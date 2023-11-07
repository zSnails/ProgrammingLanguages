package main

import "fmt"

type Socio struct {
	Numero    uint32      `json:"numero"`
	Nombre    string      `json:"nombre"`
	Direccion string      `json:"direccion"`
	Prestamos []*Prestamo `json:"-"`
}

func (s Socio) String() string {
	return fmt.Sprintf("Socio(%d, %s, %s)", s.Numero, s.Nombre, s.Direccion)
}

func NewSocio(numero uint32, nombre, direccion string) *Socio {
	return &Socio{
		Numero:    numero,
		Nombre:    nombre,
		Direccion: direccion,
		Prestamos: []*Prestamo{},
	}
}
