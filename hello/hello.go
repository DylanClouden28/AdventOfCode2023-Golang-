package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	//This is a comment
	/* This is a multi line
	comment
	*/
	fmt.Println(quote.Go())

	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	/////////////////////////////

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	println(c)
}
