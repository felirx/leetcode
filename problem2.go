package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func ToArray(list *ListNode) []int {
    if list == nil {
        panic("No mate")
    }

    arr := make([]int, 0)
    arr = toArrayRec(list, arr)
    return arr
}

func toArrayRec(it *ListNode, arr []int) []int {
    if it == nil {
        return arr
    }
    arr = append(arr, it.Val)
    return toArrayRec(it.Next, arr)
}

func Problem2() {
    l1 := &ListNode{Val: 2, Next:
    &ListNode{Val: 4, Next:
    &ListNode{Val: 3, Next: nil,},},}
    l2 := &ListNode{ Val: 5, Next:
    &ListNode{Val: 6, Next:
    &ListNode{Val: 4, Next: nil,},},}
    fmt.Println(ToArray(addTwoNumbers(l1, l2)))

    l3 := &ListNode{Val: 9, Next:
    &ListNode{Val: 9, Next:
    &ListNode{Val: 9, Next:
    &ListNode{Val: 9, Next:
    &ListNode{Val: 9, Next:
    &ListNode{Val: 9, Next:
    &ListNode{Val: 9, Next: nil },},},},},},}

    l4 := &ListNode{Val: 9, Next:
    &ListNode{Val: 9, Next:
    &ListNode{Val: 9, Next:
    &ListNode{Val: 9, Next: nil },},},}
    fmt.Println(ToArray(addTwoNumbers(l3, l4)))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var retVal *ListNode
    var l1It *ListNode = l1
    var l2It *ListNode = l2
    var retIt *ListNode
    carry := 0
    for {
        sum := carry
        if l1It == nil && l2It == nil {
            for carry > 0 {
                retIt.Next = &ListNode{Val: carry % 10, Next: nil}
                carry = carry / 10
                retIt = retIt.Next
            }
            return retVal
        }
        
        if l1It != nil {
            sum += l1It.Val
            l1It = l1It.Next
        }
        if l2It != nil {
            sum += l2It.Val
            l2It = l2It.Next
        }
        if sum >= 10 {
            carry = sum / 10
            sum = sum % 10
        } else {
            carry = 0
        }
        if retIt == nil {
            retVal = &ListNode{Val: sum, Next: nil}
            retIt = retVal
        } else {
            retIt.Next = &ListNode{Val: sum, Next: nil}
            retIt = retIt.Next
        }
    }
}
