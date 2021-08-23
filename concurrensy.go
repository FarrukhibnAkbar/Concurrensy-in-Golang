package main

import(
	"fmt"
	"time"
)


func message() {
	fmt.Println("Goodbye world")
}


func main() {

	go message()

	time.Sleep(time.Millisecond * 5)

	fmt.Println("Hello world")
}