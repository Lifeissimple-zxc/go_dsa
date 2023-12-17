package linkedlist

import (
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

func (n Node) String() string {
	return strconv.Itoa(n.Val)
}

func (n *Node) Traverse() string {
	sb := strings.Builder{}
	for cursor := n; cursor != nil; cursor = cursor.Next {
		sb.WriteString(strconv.Itoa(cursor.Val))
		if cursor.Next != nil {
			sb.WriteString("->")
		}
	}
	return sb.String()

}

func MergeSortedListsUsingHeads(h1, h2 *Node) *Node {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	// Getting here means we actually need to do the merge work
	var hd *Node
	cur1, cur2 := h1, h2

	// Assign head of the new list
	if cur1.Val < cur2.Val {
		hd = cur1
		cur1 = cur1.Next
	} else {
		hd = cur2
		cur2 = cur2.Next
	}
	ret := hd

	for cur1 != nil || cur2 != nil {
		if cur1 != nil && cur2 == nil {
			hd.Next = cur1
			cur1 = cur1.Next
			hd = hd.Next
			continue
		} else if cur1 == nil && cur2 != nil {
			hd.Next = cur2
			cur2 = cur2.Next
			hd = hd.Next
			continue
		}

		if cur1.Val < cur2.Val {
			hd.Next = cur1
			cur1 = cur1.Next
			hd = hd.Next
		} else {
			hd.Next = cur2
			cur2 = cur2.Next
			hd = hd.Next
		}

	}
	return ret
}

func locateMiddle(node *Node) *Node {
	fast, slow := node, node
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}
	return slow
}

func copy(node *Node) *Node {
	if node == nil {
		return nil
	}
	// create a new pointer to head
	newHead := &Node{
		Val: node.Val,
	}
	originalCursor := node
	newCursor := newHead

	for originalCursor.Next != nil {
		// Traversal of original list
		originalCursor = originalCursor.Next
		newCursor.Next = &Node{
			Val: originalCursor.Val,
		}
		// Traversal of new list
		newCursor = newCursor.Next
	}

	return newHead
}

func reverse(node *Node) *Node {
	var prev *Node
	cursor := node
	for cursor != nil {
		// Saving next to a variable coz
		// things will change for cursor
		nxt := cursor.Next
		// Reverse next pointer of current node
		cursor.Next = prev
		// Move prev (needed for reverse the following nodes)
		prev = cursor
		// Move cursor to continue traversal
		cursor = nxt
	}
	return prev
}

func ReverseUsingHead(head *Node, inplace bool) *Node {
	listStart := head
	if !inplace {
		listStart = copy(head)
	}
	return reverse(listStart)
}

func compareUsingHead(head1, head2 *Node) bool {
	cursor1, cursor2 := head1, head2
	for cursor1 != nil && cursor2 != nil {
		if cursor1.Val != cursor2.Val {
			return false
		}
		cursor1, cursor2 = cursor1.Next, cursor2.Next
	}
	return true
}

// HasCycleHead is like HasCycle, but it operates using Node and not LinkedList
func HasCycleHead(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	seen := make(map[*Node]interface{})

	for cursor := head; cursor != nil; cursor = cursor.Next {
		// no next == no cycle
		if cursor.Next == nil {
			return false
		}
		_, ok := seen[cursor]
		// First time seen value
		if ok {
			return true
		}
		seen[cursor] = struct{}{}
	}
	return false
}

// HasCycleTwoPointersHead is like HasCycleTwoPointers,
// but takes Node as input
func HasCycleTwoPointersHead(head *Node) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// SimplePalindromeCheckHead is like SimplePalindromeCheck,
// but operates on LL Heads of type Node
func SimplePalindromeCheckHead(head *Node) bool {
	if head == nil {
		return false
	}
	// List of 1 len is always a palindrome
	if head.Next == nil {
		return true
	}
	var cursor *Node
	var listData []int // can be optimized by using make and then appending when cap is exceeded?

	for cursor = head; cursor != nil; cursor = cursor.Next {
		listData = append(listData, cursor.Val)
	}
	cursor = head

	for i := len(listData) - 1; i >= 0; i-- {
		if listData[i] != cursor.Val {
			return false
		}
		cursor = cursor.Next
	}
	return true
}

// FindInterSectionHeads is like FindInterSection
// but takes heads of two linked lists as inputs
func FindIntersectionHeads(l1 *Node, l2 *Node) *Node {
	// Allocate a map to store traversal results
	listData := make(map[int][]*Node)
	// Traverse l1 saving it's data to a map of int[*Node] data structure
	var cursor *Node
	for cursor = l1; cursor != nil; cursor = cursor.Next {
		arr, ok := listData[cursor.Val]
		if !ok {
			res := make([]*Node, 1)
			res[0], listData[cursor.Val] = cursor, res
			continue
		}
		listData[cursor.Val] = append(arr, cursor)
	}
	// Traverse l2 doing lookups
	for cursor = l2; cursor != nil; cursor = cursor.Next {
		nodes, ok := listData[cursor.Val]
		if !ok {
			continue
		}
		for _, node := range nodes {
			if node == cursor {
				return node
			}
		}
	}
	// Reaching this point means there is no intersection
	return nil
}
