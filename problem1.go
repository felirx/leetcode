package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func Problem1() {
    input := make([]int, 10000)
    for index := range input {
        input[index] = -1000000000 + rand.Intn(2000000000)
    }
    target := -1000000000 + rand.Intn(2000000000)
    fmt.Printf("target: %d, indices: %d\n", target, twoSum(input, target))
}

func Problem1_1() {
    fmt.Println(twoSum2([]int{2,7,11,15}, 9))
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
                return []int{ i, k }
            }
        }
    }
    return []int{}
}

type data struct {
    Original int
    Value int
}

// ???
func twoSum2(nums []int, target int) []int {
    if len(nums) < 2 {
        panic ("No mate")
    }

    length := len(nums)
    sorted := make([]data, 0, length)
    for i := 0; i < length; i++ {
        sorted = append(sorted, data{Original: i, Value: nums[i]})
    }
    sort.Slice(sorted, func(a, b int) bool {
        return sorted[a].Value < sorted[b].Value
    })
    left := 0
    right := length -1
    for left != right {
        sum := sorted[left].Value + sorted[right].Value
        if sum == target {
            return []int{sorted[left].Original, sorted[right].Original}
        } else if sum < target {
            left++
        } else {
            right--
        }
    }

    return []int{}
}
