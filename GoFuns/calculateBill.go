package main

import (  
	"fmt"
	)

//func calculateBill(price int, no int) int {  
func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}


func main() {  
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)
}

