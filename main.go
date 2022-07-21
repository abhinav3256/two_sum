package main

import "fmt"

func main() {

	nums := []int{3, 3}
	result := []int{}
	target := 6
out:
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				result = append(result, i, j)
				break out
			}
		}
	}

	fmt.Println(result)
}
