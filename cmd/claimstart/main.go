package main

import (
	"fmt"
)

var (
	stat = make(map[int]int)
)

func claimStar(n int) (num int) {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	stat[n]++

	return claimStar(n-1) + claimStar(n-2)
}

func main() {
	fmt.Println("Hello world.")

	total := claimStar(9)
	fmt.Println("total is ", total)

	for key, value := range stat {
		fmt.Printf("n is %d times %d\n ", key, value)
	}

}
