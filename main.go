package main

import (
	"fmt"
	"log"
)

func main() {

	mtx1 := [][]int{
		{0, 5, 4, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 3, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	ua := []int{0, 2, 3}

	percent, err := EvalSequence(mtx1, ua)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(percent)

}
