package main

import (
	"errors"
	"fmt"

	// XXX: profe o asistente, si no le funca revise la versión de go, que en
	// la 1.21 agregaron el paquete slices como builtin de la librería estándar
	// y ya es necesario importarlo desde golang.org/x/exp/slices
	"slices"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}

func (p producto) String() string {
	return fmt.Sprintf("producto { %s, %d, %d }", p.nombre, p.cantidad, p.precio)
}

type listaProductos []producto

const EXISTENCIA_MINIMA int = 10

func (l *listaProductos) agregar(productos ...producto) {
	for _, p := range productos {
		prod := l.buscar(func(pn *producto) bool {
			return pn.nombre == p.nombre
		})
		if prod != nil {
			prod.cantidad += p.cantidad
			prod.precio = p.precio
			continue
		}
		*l = append(*l, p)
	}

}

func (l *listaProductos) buscar(indexFunc func(p *producto) bool) *producto {
	for idx := 0; idx < len(*l); idx++ {
		p := &(*l)[idx]
		if indexFunc(p) {
			return p
		}
	}

	return nil
}

func eliminar(l listaProductos, idx int) listaProductos {
	return append(l[:idx], l[idx+1:]...)
}

func (l *listaProductos) index(p *producto) int {
	for idx := 0; idx < len(*l); idx++ {
		if &(*l)[idx] == p {
			return idx
		}
	}
	return -1
}

func (l *listaProductos) vender(prod *producto, cantidad int) error {
	if prod == nil {
		return errors.New("no se puede vender un producto nil")
	}

	if prod.cantidad-cantidad < 0 {
		return errors.New("existencias insuficientes")
	}

	prod.cantidad -= cantidad
	if prod.cantidad == 0 {
		(*l) = eliminar(*l, l.index(prod))
	}

	return nil
}

func (l listaProductos) minimos() (out listaProductos) {
	for _, producto := range l {
		if producto.cantidad < EXISTENCIA_MINIMA {
			out = append(out, producto)
		}
	}
	return
}

// NOTE: esta es la función para el punto a
func (l *listaProductos) reabastecer() {
	productos := l.minimos()
	for _, producto := range productos {
		producto.cantidad += EXISTENCIA_MINIMA - producto.cantidad
		l.agregar(producto)
	}
}

func main() {
	productos := listaProductos{}
	productos.agregar(
		producto{nombre: "arroz", cantidad: 15, precio: 1500},
		producto{nombre: "frijoles", cantidad: 15, precio: 2300},
		producto{nombre: "leche", cantidad: 15, precio: 5700},
		producto{nombre: "cafe", cantidad: 15, precio: 1700},
		producto{nombre: "arroz", cantidad: 15, precio: 2500}, // NOTE: como este ya está entonces se actualizarán sus valores
	)

	fmt.Printf("productos: %v\n", productos)

	arroz := productos.buscar(func(p *producto) bool {
		return p.nombre == "arroz"
	})

	err := productos.vender(arroz, 20)
	if err != nil {
		panic(err)
	}
	fmt.Printf("productos luego de vender arroz: %v\n", productos)

	productos.reabastecer()

	// NOTE: aquí está el punto b
	slices.SortFunc(productos, func(a producto, b producto) int {
		if a.precio < b.precio {
			return -1
		} else if b.precio > a.precio {
			return 1
		}
		return 0
	})
}
