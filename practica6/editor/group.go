package main

type Group struct {
	*Object
}

// String implements DataObject.
func NewGroup() Group {
	return Group{
		Object: &Object{
			Children: []DataObject{},
		},
	}
}
