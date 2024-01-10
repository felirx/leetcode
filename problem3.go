package main

import "fmt"

func Problem3() {
    fmt.Println(lengthOfLongestSubstring("abcabcbb"))
    fmt.Println(lengthOfLongestSubstring("bbbbb"))
    fmt.Println(lengthOfLongestSubstring("pwwkew"))

    fmt.Println(lengthOfLongestSubstring2("abcabcbb"))
    fmt.Println(lengthOfLongestSubstring2("bbbbb"))
    fmt.Println(lengthOfLongestSubstring2("pwwkew"))
    fmt.Println(lengthOfLongestSubstring2("au"))
}

// O(i*j) solution
func lengthOfLongestSubstring(s string) int {
    runes := []rune(s)
    length := len(runes)
    //lengths := make([]int, length)
    longestSeen := 0
    for i := 0; i < length; i++ {
        seen := make(map[rune]int8)
        substringLength := 0
        for j := i; j < length; j++ {
            _, ok := seen[runes[j]]
            if ok {
                // we have seen this rune
                break
            } else {
                seen[runes[j]] = 1
                substringLength++
            }
        }
        //lengths[i] = substringLength
        if substringLength > longestSeen {
            longestSeen = substringLength
        }
    }
    return longestSeen
}

// Sliding window
func lengthOfLongestSubstring2(s string) int {
    runes := []rune(s)
    length := len(runes)
    if length < 2 {
        return length
    }

    longestSeen := 0
    left, right := 0, 0
    seen := make(map[rune]int)
    for right < length {
        duplicateIndex, ok := seen[runes[right]]
        if !ok {
            seen[runes[right]] = right
            right++
            continue
        }
        dist := right - left
        if dist > longestSeen {
            longestSeen = dist
        }
        // wind left
        if !ok {
            panic("It gone")
        }
        for left <= duplicateIndex {
            delete(seen, runes[left])
            left++
        }
    }
    dist := right - left
    if dist > longestSeen {
        longestSeen = dist
    }
    return longestSeen
}
