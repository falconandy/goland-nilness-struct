package main

import (
	"fmt"
)

type A struct {
	f string
}

type B struct {
	a A
}

func main() {
	b, err := createB()
	if err != nil {
	}

	fmt.Println(b.a.f)
}

func createB() (B, error) {
	return B{}, nil
}
