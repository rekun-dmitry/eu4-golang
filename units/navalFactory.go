package main

type navalFactory struct {
}

func (a *navalFactory) makeUnit() warUnit {
	return &navalUnit{
		unit: unit{
			name:            "Carrack",
			buyCost:         10,
			maintenanceCost: 0.5,
			morale:          2.56,
			combatAbility:   1.0,
		},
		sailors:        1000,
		tradePower:     0,
		hull:           25,
		cannons:        50,
		speedTactic:    5,
		speedStrategic: 6,
	}
}
