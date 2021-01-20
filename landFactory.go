package main

type landFactory struct {
}

func (a *landFactory) makeUnit() warUnit {
	return &landUnit{
		unit: unit{
			name:            "Bardiche Infantry",
			buyCost:         10,
			maintenanceCost: 0.5,
			morale:          2.56,
			combatAbility:   1.0,
		},
		manpower: 1000,
		fire:     1,
		shock:    1,
	}
}
