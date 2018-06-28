package main

import (  
	"fmt"
)

func rectDim(length, width float64)(float64, float64) {  
	var area = length * width
	var perimeter = (length + width) * 2
	
	return area, perimeter
	//return
}


func main() {  
	area, perimeter := rectDim(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f \n", area, perimeter) 
}
