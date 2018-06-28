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
	area, _ := rectDim(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f \n", area)
}
