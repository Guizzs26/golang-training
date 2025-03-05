/*

Description:

Ask the user to enter a number (as a string) and convert this value to int and float64.
Print the converted values.

*/

package main

import (
	"fmt"
	"strconv"
)

func ex4() {
	var numAsStr string

	fmt.Print("Digite um número: ")
	fmt.Scanln(&numAsStr)

	numConvertedToInt, err := strconv.Atoi(numAsStr)
	if err != nil {
		fmt.Println("Error converting to int", err)
		return
	}

	numConvertedToFloat64, err := strconv.ParseFloat(numAsStr, 64)
	if err != nil {
		fmt.Println("Erro ao converter para float64", err)
		return
	}

	fmt.Printf("Int: %d", numConvertedToInt)
	fmt.Printf("Int: %.2f", numConvertedToFloat64)

	fmt.Printf("Numebr as int: %d\n", numConvertedToInt)
	fmt.Printf("Numebr as float64: %.2f\n", numConvertedToFloat64)
}
