package main

import (
	"errors"
	"fmt"
)

const (
	DirLeft = iota
	DirRight
)

func rotate[T any](slice []T, point int, direction int) error {
	point = point % len(slice)
	tmp := make([]T, point)
	length := len(slice)

	switch direction {
	case DirLeft:
		copy(tmp, slice[:point])
		copy(slice, slice[point:])
		copy(slice[length-point:], tmp)
	case DirRight:
		copy(tmp, slice[length-point:])
		copy(slice[point:], slice[:length-point])
		copy(slice, tmp)
	default:
		return errors.New("Unknown direction")
	}

	return nil
}

func main() {
	ages := []int{1, 2, 3, 4}
    err := rotate(ages, 1, 4)
    if err != nil {
        panic(err)
    }
	fmt.Println(ages)
}
