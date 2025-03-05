/*

Description:

Declare a decimal number (float64), convert to int and print before and after the conversion.

*/

package main

import "fmt"

func ex5() {
	num := 21.75

	fmt.Printf("Float number after convertion: %.2f", num)

	numConvertedToInt := int(num)
	fmt.Printf("Float number before convertion: %d", numConvertedToInt)
}
