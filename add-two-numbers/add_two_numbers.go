package main

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func (node *ListNode) add(data int) {
	if node != nil {
		n := &ListNode{
			Next: nil,
			Val:  data,
		}

		for node.Next != nil {
			node = node.Next
		}

		node.Next = n
	}
	node.Val = data
}

func (node *ListNode) print() {
	if node == nil {
		fmt.Println("list is emtpy")
	}

	head_node := node

	for node.Next != nil {
		fmt.Printf("%d\n", node.Val)
		node = node.Next
	}

	node = head_node
}

func add_two_number(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		Val:  0,
		Next: nil,
	}
	head := result

	carry_over := 0
	left_over := 0

	end_of_l1 := false
	end_of_l2 := false

	l1_val := l1.Val
	l2_val := l2.Val

	for {
		left_over = l1_val + l2_val + carry_over

		if left_over > 9 {
			left_over %= 10
			result.Val = left_over
			carry_over = 1
		} else {
			result.Val = left_over
			carry_over = 0
		}

		if l1.Next != nil {
			l1_val = l1.Next.Val
			l1 = l1.Next
		} else {
			end_of_l1 = true
			l1_val = 0
		}

		if l2.Next != nil {
			l2_val = l2.Next.Val
			l2 = l2.Next
		} else {
			end_of_l2 = true
			l2_val = 0
		}

		if end_of_l1 && end_of_l2 && carry_over == 0 {
			return head
		}

		result.Next = &ListNode{Val: 0, Next: nil}
		result = result.Next
	}
}

func main() {

	node := ListNode{}
	node.add(2)
	node.add(4)
	node.add(3)
	node2 := ListNode{}
	node2.add(5)
	node2.add(6)
	node2.add(4)

	result := add_two_number(&node, &node2)
	result.print()
}
