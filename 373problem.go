package main

import (
	"fmt"
	"sort"
)

//Given a list of integers L, find the maximum length of a sequence of consecutive numbers that can be formed using elements from L.
//For example, given L = [5, 2, 99, 3, 4, 1, 100], return 5 as we can build a sequence [1, 2, 3, 4, 5] which has length 5.

func main() {
	fmt.Println(consecutive([]int{5, 2, 99, 3, 4, 1, 100, 98, 97, 96, 95}))
}
/*
func consecutive(arr []int) int {

	sort.Ints(arr)
	diff := -1

	var ar2 []int

	for i := 0; i < len(arr)-1; i++ {
		if (arr[i] - arr[i+1]) == diff {
			ar2 = append(ar2, arr[i])
		} else {
			continue
		}
	}
	return len(ar2)

}
*/
func consecutive(arr []int) int {

	sort.Ints(arr)
	fmt.Println(arr)

	temp := make(map[int]int)
	res := 0
	for i := 1; i < len(arr); i++ {
		fmt.Println(i, arr[i])

		r := arr[i] - arr[i-1]
		if r == 1 {
			temp[arr[i]] = 0
			temp[arr[i-1]] = 0
		} else {
			//fmt.Println(temp)
			if len(temp) > res {
				res = len(temp)
			}
			temp = make(map[int]int)
		}

	}

	fmt.Println(temp)
	if len(temp) > res {
		res = len(temp)
	}

	return res

}