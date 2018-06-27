package main

import (  
	"fmt"
	"unsafe"
	)

func main() {  
	var a int = 89
	b := 95
	var c byte = 53
	var d rune =0


	fmt.Println("value of (a int = 89) is", a, "and (b :=95) is", b, "and (var c byte = 53) is", c, "and(var d rune =0) is",d)
	fmt.Printf("\ntype of a is %T, size of a is %d", a, unsafe.Sizeof(a)) //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d \n", b, unsafe.Sizeof(b)) //type and size of b

	fmt.Printf("\ntype of c is %T, size of c is %d \n", c, unsafe.Sizeof(c)) // type and size of c
	fmt.Printf("\ntype of d is %T, size of d is %d \n", d, unsafe.Sizeof(d)) // type and size of d
	
	}
