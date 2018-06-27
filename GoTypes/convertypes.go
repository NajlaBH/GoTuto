package main

import (  
	"fmt"
    )

    
func main() {  
	i := 55      //int
	j := 67.8    //float64
	
	//sum := i + j //int + float64 not allowed
	sum := i + int(j)
	fmt.Println("Declaration i:=55, j:=67.8")
	fmt.Println("sum := i + int(j) =>",sum)

	t := 10
	var m float64 = float64(t) //this statement will not work without explicit conversion
	fmt.Println("Declaration t:=10, var m float64 = float64(t)")
	fmt.Println("t=", t, "-", "m=",m)
}
