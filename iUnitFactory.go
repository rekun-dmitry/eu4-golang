package main

import "fmt"

type iUnitFactory interface {
	makeUnit() iUnit
}

func getUnitFactory(unitType string) (iUnitFactory, error) {
	if unitType == "infantry" {
		return &infantry{}, nil
	}

	if unitType == "cavalry" {
		return &cavalry{}, nil
	}

	return nil, fmt.Errorf("Wrong unit type passed")
}
