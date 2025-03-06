/*

Description:

Ask the user to enter a letter and determine whether it is a vowel or consonant.

*/

package main

import (
	"fmt"
	"strings"
)

func ex10() {
	var letter string

	fmt.Printf("Enter a single letter (consonant or vowel): ")
	fmt.Scanln(&letter)

	letter = strings.ToLower(letter)

	if letter == "a" || letter == "e" || letter == "i" || letter == "o" || letter == "u" {
		fmt.Println("It's a vowel!")
	} else {
		fmt.Println("It's a consonant.")
	}
}
