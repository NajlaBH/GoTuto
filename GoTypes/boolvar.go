package main

import "fmt"

func main() {  
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	
	//Test 1: AND
	c := a && b
	fmt.Println("c:=a && b => c=", c)
	
	//Test2: OR
	d := a || b

	fmt.Println("d:=a || b => d=", d)

	}
