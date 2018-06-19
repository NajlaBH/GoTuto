package main

import "fmt"

func main() {  
	    var myage int // variable declaration
	        fmt.Println("declar. no initial value age is", myage)
	    
	    var yourage int = 2 // variable declaration with initial value
	    	fmt.Println("declar. with initial value age is", yourage)
		
	    var ageplay int // variable declaration
		fmt.Println("ageplay initial is ", ageplay)
	    
	    ageplay = 29 //assignment
	    	fmt.Println("ageplay first assignment is", ageplay)
		
	    ageplay = 54 //assignment
	    	fmt.Println("ageplay secibd assignement  is", ageplay)
	    
	   var inferredvar = 5 // type will be inferred

		        fmt.Println("I'm an inferred var", inferredvar)
	}
