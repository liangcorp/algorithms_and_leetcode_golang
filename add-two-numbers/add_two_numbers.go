package main

import "fmt"

type ListNode struct {
	Next *ListNode
	Var  int
}

func (node *ListNode) add(data int) {
	if node != nil {
		n := &ListNode{
			Next: nil,
			Var:  data,
		}

		for node.Next != nil {
			node = node.Next
		}

		node.Next = n
	}
	node.Var = data
}

func (node *ListNode) print() {
	if node == nil {
		fmt.Println("list is emtpy")
	}

	head_node := node

	for node.Next != nil {
		fmt.Printf("%d\n", node.Var)
		node = node.Next
	}

	node = head_node
}

func add_two_number(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{
		Next: nil,
		Var:  0,
	}
	head := result

	carry_over := 0
	left_over := 0

	for list1 != nil || list2 != nil || carry_over == 1 {
		if list1 == nil && list2 != nil {
			left_over = list2.Var + carry_over

			if left_over > 9 {
				left_over %= 10
				carry_over = 1
			} else {
				carry_over = 0
			}
			list2 = list2.Next
		} else if list1 != nil && list2 == nil {
			left_over = list1.Var + carry_over

			if left_over > 9 {
				left_over %= 10
				carry_over = 1
			} else {
				carry_over = 0
			}
			list1 = list1.Next
		} else if list1 == nil && list2 == nil && carry_over == 1 {
			left_over = carry_over
			carry_over = 0
		} else {
			left_over = list1.Var + list2.Var + carry_over

			if left_over > 9 {
				left_over %= 10
				carry_over = 1
			} else {
				carry_over = 0
			}

			list1 = list1.Next
			list2 = list2.Next
		}

		node := &ListNode{
			Next: nil,
			Var:  left_over,
		}

		result = node
		if list1 != nil || list2 != nil || carry_over == 1 {
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
	node2.add(8)

	result := add_two_number(&node, &node2)
	result.print()
}
