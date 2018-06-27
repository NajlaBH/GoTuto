package main

import (  
	"fmt"
	)
func main() {  
	a, b := 5.67, 8.97
	
	fmt.Printf("Declaration: a,b:= 5.67, 8.97\n")
	fmt.Printf("type of a %T b %T\n", a, b)
		        
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)
	
	no1, no2 := 56, 89
	fmt.Printf("Declaration nO1,nO2 := 56, 89 \n")
	fmt.Printf("type of no1 %T no2 %T\n", no1, no2)
	fmt.Println("sum", no1+no2, "diff", no1-no2)
}

