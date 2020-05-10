package main

// ListNode represents a node in a linked list.
type ListNode struct {
	val  int
	next *ListNode
}

func newNode(num int) ListNode {
	return ListNode{val: num, next: nil}
}

//           i
// nums = [1,2,3,4, 5]
//             j

// linkedList = []ListNode{}

// node = linkedList[i]
// node.val = 1
// node.next = linkedList[j]

// node = linkedList[i]
// node.val = 2
// node.next = linkedList[j]

// node = linkedList[i]
// node.val = 3
// node.next = linkedList[j]

// node = linkedList[i]
// node.val = 4
// node.next = linkedList[j]

// node = linkedList[i]
// node.val = 5
// node.next = nil

func generateLinkedList(nums ...int) []ListNode {
	linkedList := make([]ListNode, len(nums))
	j := 1

	for i, v := range nums {
		node := &linkedList[i]
		node.val = v
		if j >= len(nums) {
			node.next = nil
		} else {
			node.next = &linkedList[j]
			j++
		}
	}
	return linkedList
}

func traverseList(node *ListNode) {
	for node != nil {
		node = node.next
	}
}
