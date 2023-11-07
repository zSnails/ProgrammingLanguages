package main

import (
	"errors"
	"fmt"
)

type Prestamos []*Prestamo

func (p *Prestamos) CrearPrestamo(libro *Libro, socio *Socio) error {
	if !libro.Disponible && libro.Prestado {
		return errors.New("Este libro ya está prestado")
	}
	prestamo := NewPrestamo(libro.Codigo, socio.Numero)

	libro.Prestado = true
	libro.Disponible = false

	*p = append(*p, prestamo)
	socio.Prestamos = append(socio.Prestamos, prestamo)
	return nil
}

func filter[T any](s []T, f func(T) bool) (result []T) {
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return
}

func (s *Socios) Socios3Libros() []*Socio {
	return filter(socios, func(socio *Socio) bool {
		return len(socio.Prestamos) > 3
	})
}

type Libros []*Libro

func (l *Libros) AgregarLibros(libro ...*Libro) {
	*l = append(*l, libro...)
}

type Socios []*Socio

func (s *Socios) AgregarSocios(socio ...*Socio) {
	*s = append(*s, socio...)
}

var (
	socios    = Socios{}
	libros    = Libros{}
	prestamos = Prestamos{}
)

func main() {
	socios.AgregarSocios(
		NewSocio(1, "Aaron González", "La Tigra"),
		NewSocio(2, "Samuel González", "La Tigra"),
		NewSocio(3, "Paula González", "La Tigra"),
		NewSocio(4, "Erick González", "La Tigra"),
		NewSocio(5, "Brandon No Se", "Fortuna"),
	)

	libros.AgregarLibros(
		NewLibro(1, "La Jumpa", "Arcangel"),
		NewLibro(2, "Columbia", "Quevedo"),
		NewLibro(3, "Diablo", "Loco"),
		NewLibro(4, "Ruthless", "The Marías"),
		NewLibro(5, "Hold it together", "The Marías"),
		NewLibro(6, "Bossa No Sé", "Cuco"),
		NewLibro(7, "Cold Heart", "Elton John"),
		NewLibro(8, "Spicy", "Ty Dolla $ign"),
		NewLibro(9, "A Tu Merced", "Bad Bunny"),
	)

	// Antes de prestarlos
	fmt.Println("antes")
	for _, libro := range libros {
		fmt.Printf("%s\n", libro)
	}
	// Le prestamos 3 libros a aaron
	err := prestamos.CrearPrestamo(libros[0], socios[0])
	if err != nil {
		panic(err)
	}
	err = prestamos.CrearPrestamo(libros[1], socios[0])
	if err != nil {
		panic(err)
	}
	err = prestamos.CrearPrestamo(libros[2], socios[0])
	if err != nil {
		panic(err)
	}
	err = prestamos.CrearPrestamo(libros[3], socios[0])
	if err != nil {
		panic(err)
	}

	// Luego de prestarlos
	fmt.Println("después")
	for _, libro := range libros {
		fmt.Printf("%s\n", libro)
	}

	// Obtenemos los socios que tienen prestados más de 3 libros
	fmt.Println("\nSocios con más de 3 libros antes")
	res := socios.Socios3Libros()
	for _, socio := range res {
		fmt.Printf("%s\n", socio)
	}

	err = prestamos.CrearPrestamo(libros[4], socios[1])
	if err != nil {
		panic(err)
	}
	err = prestamos.CrearPrestamo(libros[5], socios[1])
	if err != nil {
		panic(err)
	}
	err = prestamos.CrearPrestamo(libros[6], socios[1])
	if err != nil {
		panic(err)
	}
	err = prestamos.CrearPrestamo(libros[7], socios[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("\nSocios con más de 3 libros despues")
	res = socios.Socios3Libros()
	for _, socio := range res {
		fmt.Printf("%s\n", socio)
	}

	fmt.Println("\nUltimo estado de los libros")
	for _, libro := range libros {
		fmt.Printf("%s\n", libro)
	}
}
