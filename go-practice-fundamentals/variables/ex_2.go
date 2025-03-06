/*

Description:

Declare two integer variables a and b. Calculate and print the sum, subtraction,
multiplication and division (with float64).

Example:

a = 10, b = 3
Sum: 13
Subtraction: 7
Multiplication: 30
Division: 3.33

*/

package main

import "fmt"

func ex2() {
	x := 10
	y := 20

	sum := x + y
	subtraction := x - y
	multiplication := x * y
	division := float64(x) / float64(y)

	fmt.Printf("a = %d, b = %d\n", x, y)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Subtraction: %d\n", subtraction)
	fmt.Printf("Multiplication: %d\n", multiplication)
	fmt.Printf("Division: %2.f\n", division)

	/*

		We could also have done the operations on the second argument of the Printf function, like this:

		fmt.Printf("Sum: %d", a + b)

	*/
}
