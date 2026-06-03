package main

import (
	"fmt" // This package is for like print stuff
	"math"
	"math/rand" // For random numbers
	"time"      // Its like datetime for time.Now()
)

func sums(x, y int) int {
	result := x * y
	return result
}

func add1(x, y int) (int, int) {
	x1 := x + 1
	y1 := y + 1

	return x1, y1
}

// naked return / also naming return values

func testing2(one, two int) (x, y int) {
	x = one + 2
	y = two + 3
	return
}

func main() {
	fmt.Println("Testing 123")

	fmt.Println("The time is", time.Now())

	fmt.Println("number", rand.Intn(10))

	fmt.Printf("number %g Here\n", math.Sqrt(7))

	result2 := sums(1, 2)
	fmt.Println(result2)

	x2, y2 := add1(2, 4)
	fmt.Println(x2, y2)

	fmt.Println(testing2(3, 7))

}
