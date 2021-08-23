package main

import(

	"fmt"
)

func sum(y []int, ch chan int) {
	
	var result int

	for _, i := range y {
		
		result += i
	}
	
	ch <- result
}



func main () {
	
	tunel := make(chan int)
	slice := []int{1,2,3,4,5}

	go sum(slice, tunel)

	x := <- tunel

	fmt.Println(x)
}