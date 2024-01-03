package main

import (
	"fmt"
	"math/rand"
)

func Problem1() {
    input := make([]int, 10000)
    for index := range input {
        input[index] = -1000000000 + rand.Intn(2000000000)
    }
    target := -1000000000 + rand.Intn(2000000000)
    fmt.Printf("target: %d, indices: %d\n", target, twoSum(input, target))
}
// le brute force
func twoSum(nums []int, target int) []int {
    if len(nums) < 2 {
        panic("No mate")
    }

    length := len(nums)
    for i := 0; i < length; i++ {
        for k := i + 1; k < length; k++ {
            if nums[i] + nums[k] == target {
                return []int { i, k }
            }
        }
    }
    return []int {}
}
