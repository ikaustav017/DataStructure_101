// hello.go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World! from Kaustav")
	fmt.Println("go" + "lang")
	//Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	//Booleans, with boolean operators as you’d expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//3
	var a string = "initial"
	fmt.Println(a)
	//You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)
	//Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)
	//Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)
	//The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short" in this case.
	f := "short"
	fmt.Println(f)

	const n = 500000000
	//Constant expressions perform arithmetic with arbitrary precision.
	const g = 3e20 / n
	fmt.Println(g)
	//A numeric constant has no type until it’s given one, such as by an explicit cast.
	fmt.Println(int64(g))
	//A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))

}
