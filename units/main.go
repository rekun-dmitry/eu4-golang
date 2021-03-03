package main

import "fmt"

func main() {
	landFactory, _ := getUnitFactory("land")
	navalFactory, _ := getUnitFactory("naval")

	landUnit := landFactory.makeUnit()
	navalryUnit := navalFactory.makeUnit()

	printUnitDetails(landUnit)
	printUnitDetails(navalryUnit)

}

func printUnitDetails(s warUnit) {
	fmt.Printf("Name: %s", s.getName())
	fmt.Println()
}
