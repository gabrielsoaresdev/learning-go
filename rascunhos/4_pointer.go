package main

import ("fmt")

func main() {
	x := 15
	a := &x // memory address
	*a = 5 // accessing the value through the pointer

	*a = *a**a // 25
}

