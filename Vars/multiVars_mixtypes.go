package main

import "fmt"

func main() {  
	   //Method 1
	   //var width, height int = 100, 50 //declaring multiple variables
	   // var width, height = 100, 50 //"int" is dropped
	   //     fmt.Println("width is", width, "height is", height)

	   //Method 2
	   var width, height int
	       fmt.Println("width is", width, "height is", height)
	   width = 100
	   height = 50
	   	fmt.Println("new width is", width, "new height is ", height)
	   
	   //Example of short declaration	
	   name, age := "naveen", 29 //short hand declaration
	   	fmt.Println("name is", name, "age is", age)

	   //Example of assignment
	   a, b := 20, 30 // declare variables a and b
	   	fmt.Println("a is", a, "b is", b)
	   b, c := 40, 50 // b is already declared but c is new
	   	fmt.Println("b is", b, "c is", c)
	   b, c = 80, 90 // assign new values to already declared variables b and c
	   	fmt.Println("changed b is", b, "c is", c)
	}
