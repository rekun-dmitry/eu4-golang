package main

type warUnit interface {
	setName(name string)
	getName() string
	setBuyCost(buyCost float32)
	getBuyCost() float32
	setMaintenanceCost(maintenanceCost float32)
	getMaintenanceCost() float32
	setMorale(morale float32)
	getMorale() float32
	setCombatAbility(combatAbility float32)
	getCombatAbility() float32
}

type unit struct {
	name            string
	buyCost         float32
	maintenanceCost float32
	morale          float32
	combatAbility   float32
}

func (s *unit) setName(name string) {
	s.name = name
}

func (s *unit) getName() string {
	return s.name
}

func (s *unit) setBuyCost(buyCost float32) {
	s.buyCost = buyCost
}

func (s *unit) getBuyCost() float32 {
	return s.buyCost
}

func (s *unit) setMaintenanceCost(maintenanceCost float32) {
	s.maintenanceCost = maintenanceCost
}

func (s *unit) getMaintenanceCost() float32 {
	return s.maintenanceCost
}

func (s *unit) setMorale(morale float32) {
	s.morale = morale
}

func (s *unit) getMorale() float32 {
	return s.morale
}

func (s *unit) setCombatAbility(combatAbility float32) {
	s.combatAbility = combatAbility
}

func (s *unit) getCombatAbility() float32 {
	return s.combatAbility
}
