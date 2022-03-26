package main


import (
	"fmt"
)

var (
	result []int
	count int
)

const (
	N = 8
)

// 棋子放置在(row, column是否满足要求)
func isOk(row, column int ) bool{
	// 对角线所在的column
	leftUpColumn := column - 1
	rightUpColumn := column + 1

	for i:= row - 1 ; i >=0; i--{
		// 左上
		if leftUpColumn >=0 &&  result[i] == leftUpColumn  {
			return false
		}

		// 上方
		if result[i] == column {
			return false
		}

		// 右上
		if rightUpColumn < N && result[i] == rightUpColumn {
			return false
		}

		leftUpColumn--
		rightUpColumn++
	}

	return true
}

func PrintQueue() {
	count++
	for i:=0; i<N; i++{
		column  := result[i]
		out := [N]string{}
		for j:=0; j<N; j++ {
			out[j] = "*"
		}
		out[column] = "Q"

		for j:=0; j<N; j++ {
			fmt.Printf("%s ", out[j])
		}
		fmt.Println()
	}
	fmt.Println(result)
	fmt.Println("-----------")
}

// 放置第row行棋子
func Queen(row int ) {

	// 超过棋盘大小了
	if row == N {
		PrintQueue()
		return
	}

	// 挨个尝试棋子放在(row, column)
	for column :=0; column < N; column++ {
		if isOk(row, column){
			result[row] = column
			Queen(row + 1)
		}
	}

}

func main() {

	result = make([]int, N)

	Queen(0)

	fmt.Println("Count = ",count)

	fmt.Println("hello world")

	return
}
