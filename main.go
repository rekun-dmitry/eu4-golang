package main

import "fmt"

func main() {
	infantryFactory, _ := getUnitFactory("infantry")
	cavalryFactory, _ := getUnitFactory("cavalry")

	infantryUnit := infantryFactory.makeUnit()
	cavalryUnit := cavalryFactory.makeUnit()

	printUnitDetails(infantryUnit)
	printUnitDetails(cavalryUnit)

}

func printUnitDetails(s iUnit) {
	fmt.Printf("Name: %s", s.getName())
	fmt.Println()
}
