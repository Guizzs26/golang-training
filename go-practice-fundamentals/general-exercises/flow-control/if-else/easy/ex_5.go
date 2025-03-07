/*

Description:

Ask the user to enter their grade (0 to 10) and determine whether they passed (>= 7),
had to retake (5 to 6) or failed (< 5).

*/

package main

import "fmt"

func ex5() {
	var grade float64

	fmt.Print("Enter your grade (0-10): ")
	fmt.Scanln(&grade)

	if grade >= 7 {
		fmt.Println("Congratulations! You passed.")
	} else if grade >= 5 {
		fmt.Println("You are failing.")
	} else {
		fmt.Println("Unfortunately, you failed.")
	}
}
