package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	sumRec := 0
	for scanner.Scan() {
		massStr := scanner.Text()
		mass, err := strconv.ParseInt(massStr, 0, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v", err)
			os.Exit(1)
		}
		sum += calcFuel(int(mass))
		sumRec += calcFuelRecursive(int(mass))
	}
	fmt.Printf("part1: %v\npart2: %v\n", sum, sumRec)
}

func calcFuel(mass int) int {
	return mass/3 - 2
}

func calcFuelRecursive(mass int) int {
	fuel := mass/3 - 2
	if fuel <= 0 {
		return 0
	}

	return calcFuelRecursive(fuel) + fuel
}
