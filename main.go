package main

import (
	"fmt"
	"tamagochi/service"
)

func main() {

	var namePet string
	var animalPet string
	var choice int

	fmt.Println("Введите имя для питомца")
	fmt.Scan(&namePet)

	fmt.Println("Введите тип животного:")
	fmt.Scan(&animalPet)

	Pet1 := service.Pet{
		Name:      namePet,
		Animal:    animalPet,
		Happiness: 100,
		Energy:    20,
		Heath:     30,
		Hungry:    10,
		Age:       0,
	}

	for {
		// fmt.Printf("Имя:%v...Тип:%v...Голод:%v...Здоровье:%v...Счастье%v:... Энергия:%v... Возраст:%v...\n", Pet1.Name, Pet1.Animal, Pet1.Hungry, Pet1.Heath, Pet1.Happiness, Pet1.Energy, Pet1.Age)
		fmt.Println(Pet1)

		go Pet1.Decay()

		fmt.Println("1.Покормить.")
		fmt.Println("2.Отдохнуть.")
		fmt.Println("3.Поиграть.")

		fmt.Scan(&choice)
		go Pet1.Decay()
		if choice == 1 {
			Pet1.Feed()
		} else if choice == 2 {
			Pet1.Sleep()
		} else if choice == 3 {
			Pet1.Play()
		}
	}
}
