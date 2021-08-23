package main

import(
	"fmt"
	"math"
)

type ch chan int

func square(input int, channel ch){

	output := int(math.Sqrt(float64(input)))

	channel <- output
}

func cube(input int, channel ch){

	output := int(math.Pow(float64(input), 3))

	channel <- output
}

func main() {

	square_channel := make(ch)
	cube_channel := make(ch)

	go square(64, square_channel)
	go cube(4, cube_channel)

	x := <- square_channel
	y := <- cube_channel

	fmt.Printf("%v \n%v \n", x, y)
}