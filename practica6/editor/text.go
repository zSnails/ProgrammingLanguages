package main

type Text struct {
	Content string `json:"content"`
	*Object
}

func NewText(content string) Text {
	return Text{
		Object: &Object{
			Children: []DataObject{},
		},
		Content: content,
	}
}
