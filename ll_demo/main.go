package main

import (
	"fmt"
	"ll_demo/linkedlist"
)

func main() {
	ll := linkedlist.LinkedList{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)

	fmt.Println("List after creation", ll)

	ll.Reverse()

	fmt.Println("List after reversal", ll)

	ll.Reverse()

	fmt.Println("List after reversal2", ll)
}
