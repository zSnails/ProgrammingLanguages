package main

import (
	"errors"
	"fmt"
	"strings"
)

func makeStar(width int) (string, error) {
	if width < 0 {
		return "", errors.New("width has to be positive")
	} else if width%2 == 0 {
		return "", errors.New("width has to be odd")
	}

	star := strings.Builder{}

	for i := 1; i <= width/2+1; i++ {
		fmt.Fprint(&star, strings.Repeat("  ", width-i))
		fmt.Fprint(&star, strings.Repeat(" *", i*2-1))
		fmt.Fprintln(&star)
	}

	for i := width / 2; i > 0; i-- {
		fmt.Fprint(&star, strings.Repeat("  ", width-i))
		fmt.Fprint(&star, strings.Repeat(" *", i*2-1))
		fmt.Fprintln(&star)
	}

	return star.String(), nil
}

func main() {
	star, err := makeStar(21)
	if err != nil {
		panic(err)
	}

	fmt.Println(star)
}
