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

	okVal := ll.Get(4)
	fmt.Println("Ok val:", okVal)
	noVal := ll.Get(9999)
	fmt.Println("No val:", noVal)

	ll.Remove(3)
	ll.Remove(1)

	fmt.Println(ll)

	fmt.Println(ll.Remove(3))
}
