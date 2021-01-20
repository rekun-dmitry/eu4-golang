package main

import "fmt"

type warUnitFactory interface {
	makeUnit() warUnit
}

func getUnitFactory(unitType string) (warUnitFactory, error) {
	if unitType == "land" {
		return &landFactory{}, nil
	}

	if unitType == "naval" {
		return &navalFactory{}, nil
	}

	return nil, fmt.Errorf("Wrong unit type passed")
}
