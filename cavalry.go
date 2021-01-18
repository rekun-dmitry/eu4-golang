package main

type cavalry struct {
}

func (a *cavalry) makeUnit() iUnit {
	return &cavalryUnit{
		unit: unit{
			name: "Eastern Knights",
		},
	}
}
