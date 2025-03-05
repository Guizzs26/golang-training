/*

Declare an integer and a floating point number. Convert them to string and print the results.

*/

package main

import (
	"fmt"
	"strconv"
)

func ex6() {
	numInt := 19
	numFloat := 21.5

	numIntConvertedToString := strconv.Itoa(numInt)
	numFloatConvertedToString := strconv.FormatFloat(numFloat, 'f', 2, 64)

	fmt.Println("Int as string: " + numIntConvertedToString)
	fmt.Println("Float as string" + numFloatConvertedToString)
}
