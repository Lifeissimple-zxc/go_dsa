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
	Tail *Node
	Len  int
}

func (l *LinkedList) add(value int) {
	// We always start with creating a new node
	newNode := &Node{
		Val: value,
	}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	// Len increase happens always
	defer func() {
		l.Len++
		l.Tail = newNode
	}()
	// find last item the list
	cursor := l.Head // keeping it here because we need after the loop to update the ll
	// cursor traversal happens here
	for ; cursor.Next != nil; cursor = cursor.Next {
	}
	// update its pointer to add our new element there
	cursor.Next = newNode
}

func (l *LinkedList) Add(values ...int) {
	for _, val := range values {
		l.add(val)
	}
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

func MergeSortedLists(l1, l2 *LinkedList) *LinkedList {
	// Quick nil scenario checks
	if l1.Head == nil {
		return l2
	}
	if l2.Head == nil {
		return l1
	}

	resList := &LinkedList{}
	for cur1, cur2 := l1.Head, l2.Head; cur1 != nil || cur2 != nil; {
		// Check for nils
		if cur1 != nil && cur2 == nil {
			resList.add(cur1.Val)
			cur1 = cur1.Next
			continue
		} else if cur1 == nil && cur2 != nil {
			resList.add(cur2.Val)
			cur2 = cur2.Next
			continue
		}

		// Rest of the logic
		if cur1.Val < cur2.Val {
			resList.add(cur1.Val)
			cur1 = cur1.Next
		} else {
			resList.add(cur2.Val)
			cur2 = cur2.Next
		}
	}
	return resList
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
