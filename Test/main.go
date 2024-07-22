package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		var x, y int
		sum := carry
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum = carry + x + y
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}
	return dummy.Next

}

func printDummy(node *ListNode) {
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}

func main() {
	l1 := &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: nil}}}
	l2 := &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7, Next: &ListNode{Val: 1, Next: nil}}}}
	dummy := addTwoNumbers(l1, l2)
	printDummy(dummy)
	fmt.Println()

}
