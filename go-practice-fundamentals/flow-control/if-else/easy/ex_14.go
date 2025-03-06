/*

Description:

Ask the user to enter a temperature and choose whether to convert
from Celsius to Fahrenheit or vice versa.

*/

package main

import "fmt"

func ex14() {
	var temp float64
	var choice int

	fmt.Println("Choose the conversion:")
	fmt.Println("1 - Celsius to Fahrenheit")
	fmt.Println("2 - Fahrenheit to Celsius")

	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	fmt.Print("Enter the temperature: ")
	fmt.Scanln(&temp)

	if choice == 1 {
		fahrenheitTemp := (temp * 9 / 5) + 32
		fmt.Printf("Converted temperature: %.2f°F\n", fahrenheitTemp)
	} else if choice == 2 {
		celsiusTemp := (temp - 32) * 5 / 9
		fmt.Printf("Converted temperature: %.2f°C\n", celsiusTemp)
	} else {
		fmt.Println("Invalid option! Please enter 1 or 2.")
	}
}
