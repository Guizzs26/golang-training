/*

Description:

Declare an integer, convert to string and concatenate with the message “The number is” to print.

*/

package main

import (
	"fmt"
	"strconv"
)

func ex3() {
	num := 10
	numToString := strconv.Itoa((num))

	fmt.Println("The number is: " + numToString)
}
