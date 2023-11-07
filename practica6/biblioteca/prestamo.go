package main

import (
	"fmt"
	"time"
)

type Prestamo struct {
	Fecha       time.Time `json:"fecha"`
	CodigoLibro uint32    `json:"codigoLibro"`
	NumeroSocio uint32    `json:"numeroSocio"`
}

func (p Prestamo) String() string {
	return fmt.Sprintf("Prestamo(%s, %d, %d)", p.Fecha, p.CodigoLibro, p.NumeroSocio)
}

func NewPrestamo(codigoLibro uint32, numeroSocio uint32) *Prestamo {
	return &Prestamo{
		Fecha:       time.Now(),
		CodigoLibro: codigoLibro,
		NumeroSocio: numeroSocio,
	}
}
