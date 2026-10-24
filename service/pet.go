package service

import "fmt"

// Struct pet for game
type Pet struct {
	Name   string
	Animal string

	Happiness int
	Energy    int

	Heath  int
	Hungry int

	// Data-struct
	Age int
}

// Method For feeding'n heathing pet
func (p *Pet) Feed() {
	if p.Hungry >= 100 {
		fmt.Println("Я не хочу кушать, я наелся!")
		p.Hungry = 100
	} else {
		p.Hungry += 20
		p.Heath += 10

		if p.Hungry > 100 {
			p.Hungry = 100
		}
	}
}

// Method for sleep
func (p *Pet) Sleep() {
	if p.Energy >= 100 {
		p.Energy = 100
		fmt.Println("Я не хочу спать!")
	} else {
		p.Energy += 70
		p.Heath += 5
		p.Age++

		if p.Energy > 100 {
			p.Energy = 100
		}
	}
}

// Method for play with pet
func (p *Pet) Play() {
	p.Happiness += 15
	p.Energy -= 5
}
