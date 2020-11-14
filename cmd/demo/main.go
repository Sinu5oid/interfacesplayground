package main

import (
	"fmt"
	"github.com/Sinu5oid/interfacesplayground"
)

func main() {
	equations := make([]interfaces.OneArgumentEquation, 0)

	// define equations
	equationsSet := [][]float64{
		{5, 9, 4},
		{6, 9},
		{5},
		{-5, -9},
		{3, 25, -4},
		{0,  9, 4},
		{0},
		{},
	}

	// build equations from sets
	for _, set := range equationsSet {
		if len(set) < 1 {
			continue
		}

		var eq interfaces.OneArgumentEquation
		var err error
		if len(set) == 1 {
			eq, err = interfaces.NewOneArgumentEquation(set[0])
		} else {
			eq, err = interfaces.NewOneArgumentEquation(set[0], set[1:]...)
		}

		if err != nil {
			fmt.Println("an error occurred", err, "skipping this one")
			continue
		}
		equations = append(equations, eq)
	}

	x := 5.0

	// explain equations, compute function value at x
	for _, equation := range equations {
		fmt.Printf(
			"%s\ngo type: %#v\ny(%f) = %s\n\n",
			equation.Representative(),
			equation,
			x,
			interfaces.Representative(equation.Find(complex(x, 0))),
		)
	}
}
