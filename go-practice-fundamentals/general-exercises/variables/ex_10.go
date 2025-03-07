/*

Description:

Declare an int8 and try to assign a value greater than its maximum limit. Use a check to avoid overflow.

*/

package main

import (
	"fmt"
	"math"
)

func ex10() {
	var x int8 = math.MaxInt8
	fmt.Println("Maximum value of int8: ", x)

	x += 1
	fmt.Println("After overflow: ", x)
}

/*

Adding 1 causes an overflow, and the value “returns” to the minimum value.

*/
