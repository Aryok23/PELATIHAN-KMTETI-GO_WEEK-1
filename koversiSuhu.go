package main

import "fmt"

func main() {

	var temperatureScale int

	fmt.Println("Temperature Converter Tool")
	fmt.Println("==========================")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Celsius to Kelvin")
	fmt.Println("3. Celsius to Reamur")
	fmt.Print("Input your choose: ")

	fmt.Scan(&temperatureScale)

	switch temperatureScale {
	case 1:
		{
			var celsius float64

			fmt.Print("Input your celcius temperature : ")
			fmt.Scan(&celsius)

			fahrenheit := (celsius * 9.0 / 5.0) + 32
			fmt.Printf("Your converted temperature is %.2f Fahrenheit\n", fahrenheit)

		}
	case 2:
		{
			var celsius float64

			fmt.Print("Input your celcius temperature : ")
			fmt.Scan(&celsius)

			kelvin := celsius + 273
			fmt.Printf("Your converted temperature is %.2f Kelvin\n", kelvin)
		}
	case 3:
		{
			var celsius float64

			fmt.Print("Input your celcius temperature : ")
			fmt.Scan(&celsius)

			reamur := celsius * (4.0 / 5.0)
			fmt.Printf("Your converted temperature is %.2f Reamur\n", reamur)
		}
	default:
		{
			fmt.Println("invalid temperature scale")
		}

	}

}
