package main

import (  
	    "fmt"
	    "math"
	)


func main(){  
	a, b := 145.8, 543.8
	minab := math.Min(a, b)
	maxab := math.Max(a, b)
	trunca := math.Trunc(a)
	roundb := math.Round(b)

	fmt.Println("a",a, " b",b)
	fmt.Println("minimum value is ", minab)
	fmt.Println("maximum value is", maxab)
	fmt.Println("Trunc a value is", trunca)
	fmt.Println("Round b value is", roundb)
	}

