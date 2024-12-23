package main

import "fmt"

func Val() {
	fmt.Println("go" + "lang") // Strings concatenated
	fmt.Println("1+1 = ", 1+1) // Integers

	fmt.Println("7.0/3.0 =", 7.0/3.0)                  // Floats
	fmt.Println(!(true && false) || !(!true || false)) // Booleans
}
