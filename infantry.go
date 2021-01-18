package main

type infantry struct {
}

func (a *infantry) makeUnit() iUnit {
	return &infantryUnit{
		unit: unit{
			name: "Bardiche Infantry",
		},
	}
}
