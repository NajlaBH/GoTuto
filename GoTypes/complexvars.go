package main

import (  
	"fmt"
    )

func main() {  
	c1 := complex(5, 7)
	c2 := 8 + 27i
	
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	// Example : func complex(r, i FloatType) ComplexType  
	comp := 6 + 7i 
	fmt.Println("the value of (comp := 6 + 7i) is",comp )
}


