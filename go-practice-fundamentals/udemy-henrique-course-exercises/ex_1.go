/*

Description:

create an array with 2 integer positions and store the total sum of the list in a variable.
The variable should be printed to the console.

*/

package main

import "fmt"

func ex1() {
	numbers := [2]int{20, 80}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	fmt.Printf("The sum of the array is: %d", sum)
}

/*

We can also do this way:

sum := numbers[0] + numbers[1]

*/
