/*

Description:

Declare constants to store fixed values, such as π (3.14) and the number
of days in a week (7). Print the values.

*/

package main

import "fmt"

func main() {
	const pi = 3.14
	const daysOfWeek = 7

	fmt.Printf("Pi value: %.2f\n", pi)
	fmt.Println("Days of week: ", daysOfWeek)
}
