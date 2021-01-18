package main

type iUnit interface {
	setName(name string)
	getName() string
}

type unit struct {
	name string
}

func (s *unit) setName(name string) {
	s.name = name
}

func (s *unit) getName() string {
	return s.name
}
