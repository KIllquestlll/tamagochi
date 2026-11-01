package service

import (
	"fmt"
	"os"
	"time"
)

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

func (p Pet) String() string {
	return fmt.Sprintf("Имя:%v...Тип:%v...Голод:%v...Здоровье:%v...Счастье%v:... Энергия:%v... Возраст:%v...\n",
		p.Name, p.Animal, p.Hungry, p.Heath, p.Happiness, p.Energy, p.Age)
}

func (p *Pet) Decay() {
	p.Hungry -= 5
	time.Sleep(3 * time.Second)
	if p.Hungry <= 0 {
		fmt.Println("Ваш питомец умер!")
		os.Exit(1)
	}
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
