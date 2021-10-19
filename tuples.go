package main

import "fmt"

// tuples are ordered sequence of objects.They can have other data types and they are also used to group related data into a data structure.

// product function return the product of params x and y.
func p(x,y int) int {

	return x * y
}

// g function return params x and y after modification.
func g(l,m int) (x,y int){
	x = 2 * l
	y = 4 * m 
	return
}

func main(){
	fmt.Println(p(g(1,2)))
}