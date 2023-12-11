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

// FindInterSection locates where l1 and l2 intersect (memory address based)
// When there is no intersection, it returns a nil
func FindIntersection(l1 *LinkedList, l2 *LinkedList) *Node {
	// Allocate a map to store traversal results
	listData := make(map[int]*Node)
	// Traverse l1 saving it's data to a map of int[*Node] data structure
	var cursor *Node
	for cursor = l1.Head; cursor != nil; cursor = cursor.Next {
		listData[cursor.Val] = cursor
	}
	// Traverse l2 doing lookups
	for cursor = l2.Head; cursor != nil; cursor = cursor.Next {
		node, ok := listData[cursor.Val]
		if !ok {
			continue
		}
		if node == cursor {
			return node
		}
	}
	// Reaching this point means there is no intersection
	return nil
}

// FindInterSectionWithDuplicates is like FindInterSection, but
// supports linked lists where multiple nodes have the same value
func FindIntersectionWithDuplicates(l1 *LinkedList, l2 *LinkedList) *Node {
	// Allocate a map to store traversal results
	listData := make(map[int][]*Node)
	// Traverse l1 saving it's data to a map
	var cursor *Node
	for cursor = l1.Head; cursor != nil; cursor = cursor.Next {
		arr, ok := listData[cursor.Val]
		if !ok {
			// map has not data for this value
			res := make([]*Node, 1)
			res[0], listData[cursor.Val] = cursor, res
			continue
		}
		// Map has data, we append
		listData[cursor.Val] = append(arr, cursor)

	}
	// Traverse l2 doing lookups
	for cursor = l2.Head; cursor != nil; cursor = cursor.Next {
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

// FindInterSectionHeads is like FindInterSection
// but takes heads of two linked lists as inputs
func FindIntersectionHeads(l1 *Node, l2 *Node) *Node {
	// Allocate a map to store traversal results
	listData := make(map[int]*Node)
	// Traverse l1 saving it's data to a map of int[*Node] data structure
	var cursor *Node
	for cursor = l1; cursor != nil; cursor = cursor.Next {
		listData[cursor.Val] = cursor
	}
	// Traverse l2 doing lookups
	for cursor = l2; cursor != nil; cursor = cursor.Next {
		node, ok := listData[cursor.Val]
		if !ok {
			continue
		}
		if node == cursor {
			return node
		}
	}
	// Reaching this point means there is no intersection
	return nil
}

// FindIntersectionHeadsWithDuplicates is like FindIntersectionHeads
// supports linked lists where multiple nodes have the same value
func FindIntersectionHeadsWithDuplicates(l1 *Node, l2 *Node) *Node {
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

// SimplePalindromeCheck checks if a given LinkedList is palindrome
func SimplePalindromeCheck(l *LinkedList) bool {
	if l == nil {
		return false
	}
	// List of 1 len is always a palindrome
	if l.Head.Next == nil {
		return true
	}
	var cursor *Node
	var listData []int

	for cursor = l.Head; cursor != nil; cursor = cursor.Next {
		listData = append(listData, cursor.Val)
	}
	cursor = l.Head

	for i := len(listData) - 1; i >= 0; i-- {
		if listData[i] != cursor.Val {
			return false
		}
		cursor = cursor.Next
	}
	return true
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

// IsPalindrome implements a slow and fast pointer algo
// to determine if a linked list is a palindrome.
// Time complexity: O(n). Space complexity: O(1).
func IsPalindrome(l *LinkedList) bool {
	// Find middle of the list using 2 pointers (can be a func)
	subList := locateMiddle(l.Head)
	// Reverse second half (nneds to be a func to avoid modifying a list in place)
	subList = ReverseUsingHead(subList, false)
	// Compare reversed half with the first half
	// of the original list node by node
	return compareUsingHead(l.Head, subList)
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

// HasCycle checks if a given linked list has a cycle within it
func HasCycle(l *LinkedList) bool {
	if l.Head == nil || l == nil || l.Head.Next == nil {
		return false
	}
	seen := make(map[*Node]bool)

	for cursor := l.Head; cursor != nil; cursor = cursor.Next {
		// no next == no cycle
		if cursor.Next == nil {
			return false
		}
		_, ok := seen[cursor]
		// First time seen value
		if ok {
			return true
		}
		seen[cursor] = true
	}
	return false
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
