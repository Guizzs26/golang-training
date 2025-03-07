/*

Description:

Ask the user to enter a year and determine whether it is a leap year or not.

*/

package main

import "fmt"

func ex11() {
	var year int

	fmt.Println("Enter a year (integer number in interval 1900-2025)")
	fmt.Scanln(&year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("It's a leap year!")
	} else {
		fmt.Println("It's NOT a leap year!")
	}
}
