/*

Description:

Allow the user to enter a grade (0 to 100) and return the grade based on the following scale:

Grade Grade
90 - 100 A
80 - 89 B
70 - 79 C
60 - 69 D
0 - 59 F

*/

package main

import "fmt"

func ex8() {
	var grade int

	fmt.Printf("Enter your grade (0-100): ")
	fmt.Scanln(&grade)

	if grade >= 90 && grade <= 100 {
		fmt.Println("A")
	}

	if grade < 0 || grade > 100 {
		fmt.Println("Invalid grade! Please enter a value between 0 and 100.")
	} else if grade <= 59 && grade >= 0 {
		fmt.Println("Your grade is: F")
	} else if grade <= 69 {
		fmt.Println("Your grade is: D")
	} else if grade <= 79 {
		fmt.Println("Your grade is: C")
	} else if grade <= 89 {
		fmt.Println("Your grade is: B")
	} else {
		fmt.Println("Your grade is: A")
	}
}
