package main

import (
	_ "euler/challenge01"
	_ "euler/challenge02"
	_ "euler/challenge03"
	_ "euler/challenge04"
	_ "euler/challenge05"
	"euler/challenge06"
	"fmt"
)

func main() {
	result := challenge06.SumOfDifference(100)
	fmt.Printf("The answer is: %v\n", result)
}
