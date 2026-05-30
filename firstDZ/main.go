package main

import (
	"fmt"
	"math/rand/v2"
)

// final result
func main() {
	rCh := RandomSlice()
	outCh := Squared(rCh)

	result := make([]int, 0, 10)

	for num := range outCh {
		result = append(result, num)
	}

	for _, n := range result {
		fmt.Println(n)
	}
}

func RandomSlice() <-chan int {
	randomCh := make(chan int)
	go func() {
		s := make([]int, 0, 10)
		for range 10 {
			number := rand.IntN(101)
			s = append(s, number)
		}
		for _, num := range s {
			randomCh <- num
		}
		close(randomCh)
	}()
	return randomCh
}

func Squared(random <-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		for num := range random {
			squredN := num * num
			outCh <- squredN
		}
		close(outCh)
	}()
	return outCh
}
