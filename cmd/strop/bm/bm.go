package main

import (
	"fmt"
)

func GenGoodSuffix(str string, suffix[]int, prefix[]bool, m int) {
	for i:=0; i< m-1; i++ {
		k := 0
		j := i
		for ;j>=0 && str[j]==str[m-k-1]; {
			j--
			k++
			suffix[k] = j + 1
		}

		if j == -1 {
			prefix[k] = true
		}

	}

}

func main() {

	fmt.Println("hello world")

	str  := "cabbabcab"


	suffix := make([]int, len(str))
	prefix := make([]bool, len(str))

	GenGoodSuffix(str, suffix, prefix, len(str))

	fmt.Println("orignal str : ", str)

	for index, value := range suffix {
		fmt.Printf("%8s suffix[%d] = %d prefix[%d] = %t \n",
			str[len(str)-index:], index, value, index, prefix[index])
	}

	return
}