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

	for l1 != nil || l2 != nil || carry_over == 1 {
		if l1 == nil && l2 != nil {
			left_over = l2.Val + carry_over

			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			left_over = l1.Val + carry_over

			l1 = l1.Next
		} else if l1 == nil && l2 == nil && carry_over == 1 {
			left_over = carry_over
			carry_over = 0
		} else if l1 != nil && l2 != nil {
			left_over = l1.Val + l2.Val + carry_over

            // fmt.Printf("%d\t%d\n", list1.Var, list2.Var)
			l1 = l1.Next
			l2 = l2.Next
		}

		if left_over > 9 {
			left_over %= 10
			carry_over = 1
		} else {
			carry_over = 0
		}

		if l1 != nil || l2 != nil || carry_over == 1 {
			result.Val = left_over
			result.Next = &ListNode{Val: 0, Next: nil}
			result = result.Next
		}
	}
	return head
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
