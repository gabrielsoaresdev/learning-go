package main

import ("fmt")
         //x float64, y float64
func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	//It is necessary to use all declared variables
	var num1 float64 = 5.6
	var num2 float64 = 9.7

	/*
	var num1, num2 float64 = 5.6, 9.5

	num1, num2 := 5.6, 9.5
	*/

	//The type is defined after compilation
	w1, w2 := "Hey", "there"

	fmt.Println(multiple(w1, w2))

	fmt.Println(add(num1, num2))

	//Convert type
	var a int = 62
	var b float64 = float64(a)

	x := a
}

