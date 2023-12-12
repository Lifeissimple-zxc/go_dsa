package array

import "fmt"

// SimpleMaxProduct is a naive implementation of maxProduct algo
func SimpleMaxProduct(nums []int) int {
	// max1 & max2 store max values
	// maxIx stores indeces to max values in the array
	var max1, max2, maxIx1, maxIx2 int
	for cur1, cur2 := 0, -1; cur2 < len(nums); {
		// first max
		if nums[cur1] > max1 {
			max1 = nums[cur1]
			maxIx1 = cur1
		}
		// move first cursor if possible
		if cur1+1 <= len(nums)-1 {
			cur1++
		}
		// [2,2,1,8,1,5,4,5,2,10,3,6,5,2,3] fails now
		if cur2 >= 0 && nums[cur2] > max2 && cur2 != maxIx1 {
			max2 = nums[cur2]
			maxIx2 = cur2
		}
		cur2++
		fmt.Println(nums[maxIx1], nums[maxIx2])
	}
	return (nums[maxIx1] - 1) * (nums[maxIx2] - 1)

}

// TODO implement a simpler approach, it
