// implements SmallestInfiniteSet for https://leetcode.com/problems/smallest-number-in-infinite-set
package heap

type SmallestInfiniteSet struct {
	s        []int            // minheap lives here
	vals     map[int]struct{} // needed to make AddBack work in O(1)
	smallest int              // default value in case s has no data
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		s:        make([]int, 0, 1000),
		vals:     make(map[int]struct{}),
		smallest: 1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.s) == 0 {
		this.smallest++          // because we pop, our min changes
		return this.smallest - 1 // -1 accounts for the increment above
	}
	root := this.s[0] // this is the min value because we have a min heap
	// re-arrange the heap to maintain its property
	l := len(this.s) - 1
	this.s[0] = this.s[l]
	this.s = this.s[:l] // this accounts for moving lth elem to root node
	// maintain the property by heaping this down
	this.heapDown(0)
	// remove from duplicate set
	delete(this.vals, root)
	return root
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	// we don't add nums that are higher than smallest
	// because calling PopSmallest will eventually reach them anyway
	if num >= this.smallest {
		return
	}
	// we also don't add whatever is in the vals set because that's a requirement
	if _, ok := this.vals[num]; ok {
		return
	}
	this.push(num)
	this.vals[num] = struct{}{}
}

// helpers needed for exposed methods
func (this *SmallestInfiniteSet) push(num int) {
	this.s = append(this.s, num)
	this.heapUp(len(this.s) - 1)
}

func (this *SmallestInfiniteSet) heapUp(i int) {
	if this.s[getParent(i)] <= this.s[i] {
		return
	}
	this.swap(getParent(i), i)
	this.heapUp(getParent(i))
}

func (this *SmallestInfiniteSet) heapDown(i int) {
	lastIndex := len(this.s) - 1
	left, right := leftChild(i), rightChild(i)
	childToCompare := right // assume right is smaller

	if left > lastIndex {
		return
	}
	if left == lastIndex || this.s[left] < this.s[right] {
		// update our assumption of left is the only child OR smaller than right
		childToCompare = left
	}
	// base case check before recursive call
	if this.s[childToCompare] >= this.s[i] {
		// no need to swap if child is greater or eq to s[i]
		return
	}
	this.swap(i, childToCompare)
	this.heapDown(childToCompare)
}

func (this *SmallestInfiniteSet) swap(x, y int) {
	this.s[x], this.s[y] = this.s[y], this.s[x]
}

func getParent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return i*2 + 1
}

func rightChild(i int) int {
	return i*2 + 2
}
