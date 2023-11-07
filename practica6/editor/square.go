package main

import (
	"fmt"
	"strings"
)

type Square struct {
	Shape string `json:"shape"`
	*Object
}

func NewSquare(side int) Square {
	w := strings.Builder{}

	for idx := 0; idx < side; idx++ {
		fmt.Fprintf(&w, "%s\n", strings.Repeat("*", side*2))
	}

	return Square{
		Shape: w.String(),
		Object: &Object{
			Children: []DataObject{},
		},
	}
}
