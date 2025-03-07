/*

Description:

Given a slice with the items “2, 8, 3, 10, 5, 4, 7, 9, 1” ranging from 1 to 10, add them up
in two variables, the first number from 1 to 5 and the second from 6 to 10. Print out the two results.

*/

package main

import "fmt"

func ex2() {
	numbers := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	var sum1, sum2 int
	for _, number := range numbers {
		if number >= 1 && number <= 5 {
			sum1 += number
		} else if number >= 6 && number <= 10 {
			sum2 += number
		}
	}

	fmt.Println("Soma dos números de 1 a 5:", sum1)
	fmt.Println("Soma dos números de 6 a 10:", sum2)
}
