/*

Description:

Declare an int and a float64. Perform mathematical operations between them
(addition, subtraction, multiplication, division) and print the results.

*/

package main

import "fmt"

func ex9() {
	var intNumber int = 100
	var float64Number float64 = 3.14

	var additionOperation float64 = float64(intNumber) + float64Number
	var subtractionOperation float64 = float64(intNumber) - float64Number
	var multiplicationOperation float64 = float64(intNumber) * float64Number
	var divisionOperation float64 = float64(intNumber) / float64Number

	fmt.Printf("Sum: %.2f", additionOperation)
	fmt.Printf("Subtraction: %.2f", subtractionOperation)
	fmt.Printf("Multiplication: %.2f", multiplicationOperation)
	fmt.Printf("Division: %.2f", divisionOperation)
}
