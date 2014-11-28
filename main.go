package main

import (
	_ "euler/challenge01"
	_ "euler/challenge02"
	_ "euler/challenge03"
	_ "euler/challenge04"
	"euler/challenge05"
	"fmt"
)

func main() {
	result := challenge05.SmallestMultiple(20)
	fmt.Printf("The answer is: %v\n", result)
}
