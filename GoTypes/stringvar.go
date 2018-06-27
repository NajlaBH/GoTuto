package main

import (  
	    "fmt"
    )


func main() {  
	first := "Najla"
	last := "Bioinfo"
	tag := "Go2018"
	name := first +" "+ last
	fmt.Println("My name is", name, "-",tag[:2],":",tag[2:])
}
