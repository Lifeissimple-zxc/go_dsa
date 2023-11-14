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

type LinkedList struct {
	Head *Node
	Len  int
}

func (l *LinkedList) Add(value int) {
	// We always start with creating a new node
	newNode := &Node{
		Val: value,
	}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	// find last item the list
	cursor := l.Head // keeping it here because we need after the loop to update the ll
	// cursor traversal happens here
	for ; cursor.Next != nil; cursor = cursor.Next {
	}
	// update its pointer to add our new element there
	cursor.Next = newNode
}

func (l *LinkedList) Get(value int) *Node {
	// Iterate till we find our value
	for cursor := l.Head; cursor != nil; cursor = cursor.Next {
		if cursor.Val == value {
			return cursor
		}
	}
	// nil for cases when nothing was found
	return nil
}

// Remove removes node with the value from the linked list and returns it.
// Nil is returned when no element was located in the linked list.
func (l *LinkedList) Remove(value int) bool {
	var prev *Node
	for cursor := l.Head; cursor != nil; cursor = cursor.Next {
		if cursor.Val != value {
			// Not a matching value, just shift the cursor
			prev = cursor
			continue
		}
		// Handle head separateluy
		if l.Head.Val == cursor.Val {
			l.Head = cursor.Next
			return true
		}
		// Delete node with matching value by linking
		// The node preceeding it to the node following it
		prev.Next = cursor.Next
		return true
	}
	return false

}

func (l LinkedList) String() string {
	sb := strings.Builder{}
	for cursor := l.Head; cursor != nil; cursor = cursor.Next {
		sb.WriteString(strconv.Itoa(cursor.Val))
		if cursor.Next != nil {
			sb.WriteString("->")
		}
	}
	return sb.String()
}

func (l *LinkedList) Reverse() {
	if l.Head == nil || l == nil {
		return
	}
	var prev *Node
	current := l.Head
	next := l.Head.Next

	for next != nil {
		current.Next = prev // reverse current node
		prev = current      // update prev for next iteration
		current = next      // move to the next node
		next = next.Next    // update next
		current.Next = prev
	}
	l.Head = current
}

func (l *LinkedList) Reverse2() {
	if l.Head == nil || l == nil {
		return
	}
	// Less aux variables
	var reversedHead *Node // ala previous

	for l.Head != nil {
		next := l.Head.Next        // Save next for traversal
		l.Head.Next = reversedHead // Reverse node's Next
		reversedHead = l.Head
		l.Head = next
	}

	l.Head = reversedHead
}

func (l *LinkedList) ReverseRecursive() {
	if l.Head == nil || l == nil {
		return
	}

	var prev *Node

	var reversal func(l *LinkedList)
	reversal = func(l *LinkedList) {
		if l.Head.Next == nil {
			l.Head.Next = prev
			return
		}
		l.Head.Next, prev, l.Head = prev, l.Head, l.Head.Next
		reversal(l)
	}
	reversal(l)
}
