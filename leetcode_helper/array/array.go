package array

import "fmt"

// SimpleMaxProduct is a naive implementation of maxProduct algo
func SimpleMaxProduct(nums []int) int {
	// store max indexes here
	var maxIx1, maxIx2 int
	for cur1, cur2 := 0, -1; cur2 < len(nums); {
		// first max
		if nums[cur1] > nums[maxIx1] {
			maxIx1 = cur1
		}
		// move first cursor if possible
		if cur1+1 <= len(nums)-1 {
			cur1++
		}
		if cur2 >= 0 && nums[cur2] > nums[maxIx2] && cur2 != maxIx1 {
			maxIx2 = cur2
		}
		cur2++
		fmt.Println(nums[maxIx1], nums[maxIx2])
	}
	return (nums[maxIx1] - 1) * (nums[maxIx2] - 1)

}
