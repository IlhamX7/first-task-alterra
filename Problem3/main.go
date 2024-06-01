package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)

	totalHeight := 0
	j := 0

	for i := 0; i < len(dragonHead); i++ {
		for j < len(knightHeight) && knightHeight[j] < dragonHead[i] {
			j++
		}
		if j == len(knightHeight) {
			fmt.Println("knight fall")
			return
		}
		totalHeight += knightHeight[j]
		j++
	}
	fmt.Println(totalHeight)
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})
	DragonOfLoowater([]int{5, 10}, []int{5})
}
