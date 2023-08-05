package main

import (
	"errors"
	"fmt"
)

type Sellable interface {
	SetId(int)
	GetId() int
	GetPrice() int
	GetStock() int
	SetStock(int)
	GetName() string
	HasStock() bool
	fmt.Stringer
}

type Store struct {
	inventory []Sellable
}

// TODO: boundary checking
func (s *Store) FindById(id int) Sellable {
	return s.inventory[id]
}

func (s *Store) Sell(element Sellable) error {
	if !element.HasStock() {
		return errors.New(fmt.Sprintf("out of stock for '%s'", element.GetName()))
	}

	// NOTE: no hay necesidad para un boundary check aquí porque la id siempre existe
	s.inventory[element.GetId()].SetStock(element.GetStock() - 1)
	return nil
}

func (s *Store) Add(elements ...Sellable) {
	for _, element := range elements {
		element.SetId(len(s.inventory))
		s.inventory = append(s.inventory, element)
	}
}

func (s *Store) FindByName(name string) (Sellable, error) {
	for _, element := range s.inventory {
		if element.GetName() == name {
			return element, nil
		}
	}

	return nil, errors.New("that element does not exist")
}
