package main

import (
	"errors"
)

type Object struct {
	Children []DataObject `json:"children"`
}

// AddChildren implements Object.
func (t *Object) AddChildren(n ...DataObject) {
	t.Children = append(t.Children, n...)
}

// GetChild implements Object.
func (t Object) GetChild(idx int) (DataObject, error) {
	if len(t.Children) <= idx {
		return nil, errors.New("Out of bounds")
	}

	return t.Children[idx], nil
}

// GetChildren implements Object.
func (t Object) GetChildren() []DataObject {
	return t.Children
}

type DataObject interface {
	AddChildren(...DataObject)
	GetChildren() []DataObject
	GetChild(int) (DataObject, error)
}
