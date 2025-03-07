/*

Description:

Given a slice containing the numbers from 1 to 10 in random order: “2, 8, 3, 10, 5, 4, 7, 9, 1”,
your task is to:

Create two new slices:

One slice should contain the numbers between 1 and 5 (inclusive).
The second slice should contain the numbers between 6 and 10 (inclusive).
Sum the elements of each slice and print:

The contents of both slices.
The sum of the numbers in the first slice (1 to 5).
The sum of the numbers in the second slice (6 to 10).

*/

package main

import "fmt"

func ex3() {
	numbers := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	var slice1, slice2 []int
	for _, number := range numbers {
		if number >= 1 && number <= 5 {
			slice1 = append(slice1, number)
		} else if number >= 6 && number <= 10 {
			slice2 = append(slice2, number)
		}
	}

	var sum1, sum2 int
	for _, number := range slice1 {
		sum1 += number
	}

	for _, number := range slice2 {
		sum2 += number
	}

	fmt.Println("Slice 1 (1 a 5):", slice1)
	fmt.Println("Soma dos números de 1 a 5:", sum1)
	fmt.Println("Slice 2 (6 a 10):", slice2)
	fmt.Println("Soma dos números de 6 a 10:", sum2)
}
