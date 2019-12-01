package main

import (
	"fmt"
	"testing"
)

func TestCalcFuel(t *testing.T) {
	testCases := []struct {
		mass int
		fuel int
	}{
		{mass: 12, fuel: 2},
		{mass: 14, fuel: 2},
		{mass: 1969, fuel: 654},
		{mass: 100756, fuel: 33583},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v", tC), func(t *testing.T) {
			fuel := calcFuel(tC.mass)
			if tC.fuel != fuel {
				t.Errorf("error! got %v, expected %v", fuel, tC.fuel)
			}
		})
	}
}

func TestCalcFuelRecursive(t *testing.T) {
	testCases := []struct {
		mass int
		fuel int
	}{
		{mass: 14, fuel: 2},
		{mass: 1969, fuel: 966},
		{mass: 100756, fuel: 50346},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v", tC), func(t *testing.T) {
			fuel := calcFuelRecursive(tC.mass)
			if tC.fuel != fuel {
				t.Errorf("error! got %v, expected %v", fuel, tC.fuel)
			}
		})
	}
}
