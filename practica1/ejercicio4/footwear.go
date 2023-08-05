package main

import (
	"fmt"
)

type Footwear struct {
	id    int
	brand string
	model string
	price int
	size  int
	stock int
}

func (f Footwear) String() string {
	return fmt.Sprintf("footear { id = %d, brand = %s, model = %s, price = %d, size = %d, stock = %d }", f.id, f.brand, f.model, f.price, f.size, f.stock)
}

func (f Footwear) GetPrice() int {
	return f.price
}

func (f *Footwear) SetStock(stock int) {
	f.stock = stock
}

func (f Footwear) GetName() string {
	return fmt.Sprintf("%s %s", f.brand, f.model)
}

func (f Footwear) GetStock() int {
	return f.stock
}

func (f Footwear) HasStock() bool {
	return f.stock > 0
}

func (f Footwear) GetId() int {
	return f.id
}

func (f *Footwear) SetId(id int) {
	f.id = id
}

func clamp(min, max, val int) int {
	if val < min {
		return min
	} else if val > max {
		return max
	}

	return val
}

func NewFootwear(brand, model string, price, size, stock int) *Footwear {

	size = clamp(34, 44, size)

	return &Footwear{
		brand: brand,
		model: model,
		price: price,
		size:  size,
		stock: stock,
	}
}
