package main

import (  
	"fmt"
	"math"
	"unsafe"
)

func main() {  
	fmt.Println("Hello, playground")
	var a = math.Sqrt(4) //allowed
	fmt.Print("a = math.Sqrt(4) =>",a,"\n")
	//const b = math.Sqrt(4) //not allowed

	var name = "Sam"
	fmt.Printf("type %T size %d value %v\n", name, unsafe.Sizeof(name), name)
}

