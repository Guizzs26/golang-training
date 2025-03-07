/*

Description:

Ask the user to enter the price of a product and a discount percentage.
Calculate and display the final discounted amount.

*/

package main

import "fmt"

func ex13() {
	var price, discount float64

	fmt.Print("Enter the product price: ")
	fmt.Scanln(&price)

	fmt.Print("Enter the discount: (0-100%) ")
	fmt.Scanln(&discount)

	if discount < 0 || discount > 100 {
		fmt.Println("Invalid discount! Please enter a value between 0 and 100.")
		return
	}

	finalPrice := price * (1 - discount/100)
	fmt.Printf("Final price after discount: $%.2f\n", finalPrice)
}
