/*

Description:

Declare variables of types int8, int16, int32, int64, float32 and float64. Use fmt.Printf with %T to
print the type of each variable.

*/

package main

import "fmt"

func ex7() {
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807
	var e float32 = 3.14
	var f float64 = 3.141592653589793

	fmt.Printf("Datatype of a: %T\n", a)
	fmt.Printf("Datatype of b: %T\n", b)
	fmt.Printf("Datatype of c: %T\n", c)
	fmt.Printf("Datatype of d: %T\n", d)
	fmt.Printf("Datatype of e: %T\n", e)
	fmt.Printf("Datatype of f: %T\n", f)
}
