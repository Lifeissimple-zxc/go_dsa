package array

/*
https://leetcode.com/problems/kids-with-the-greatest-number-of-candies
There are n kids with candies. You are given an integer array candies,
where each candies[i] represents the number of candies the ith kid has, and an integer extraCandies,
denoting the number of extra candies that you have.

Return a boolean array result of length n, where result[i] is true if,
after giving the ith kid all the extraCandies,
they will have the greatest number of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.

Time Complexity: O(2n), linear
Space Compexity: O(n)
*/
func KidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies)) // linear space complexity bc this array grows as input does
	// Special case of 1 kid
	if len(candies) == 1 {
		res[0] = true
		return res
	}
	// General case with many kids
	// Locate the max
	maxCandies := findMax(candies) // linear
	// Compute result, also linear
	for i, candies := range candies {
		res[i] = (candies + extraCandies) >= maxCandies
	}
	return res
}

/*
https://leetcode.com/problems/can-place-flowers
You have a long flowerbed in which some of the plots are planted, and some are not.
However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty,
and an integer n, return true if n new flowers can be planted in the flowerbed
without violating the no-adjacent-flowers rule and false otherwise.

Example 1:
Input: flowerbed = [1,0,0,0,1], n = 1
Output: true

Example 2:
Input: flowerbed = [1,0,0,0,1], n = 2
Output: false


Constraints:
1 <= flowerbed.length <= 2 * 104
flowerbed[i] is 0 or 1.
There are no two adjacent flowers in flowerbed.
0 <= n <= flowerbed.length

Time Complexity: O(n), linear
Space Compexity: O(1), does not grow with input at all
*/
func CanPlaceFlowers(flowerbed []int, n int) bool {
	for c := 0; c < len(flowerbed) && n > 0; {
		if flowerbed[c] == 1 {
			c++
		} else if (c == 0 || flowerbed[c-1] == 0) &&
			(c == len(flowerbed)-1 || flowerbed[c+1] == 0) {
			c++
			n--
		}
		c++
	}
	return n <= 0
}

func ReverseVowels(s string) string {
	res := []byte(s)
	if len(res) == 1 {
		return s
	}
	vowels := map[string]struct{}{
		"a": {},
		"A": {},
		"e": {},
		"i": {},
		"E": {},
		"I": {},
		"u": {},
		"U": {},
		"o": {},
		"O": {},
	}
	// Traverse string from both ends
	// until either cursor crosses the middle
	for l, r := 0, len(res)-1; l < r; {
		_, lv := vowels[string(res[l])]
		_, rv := vowels[string(res[r])]
		// left and right are vowels
		if lv && rv {
			res[l], res[r] = res[r], res[l]
			l++
			r--
		} else if lv {
			// left is vowel, right is not, so we move just right
			r--
		} else if rv {
			// right is vowel, left is not, so we move just left
			l++
		} else {
			// both are not vowels, move both
			l++
			r--
		}
	}
	return string(res)
}

func isVowel(chr byte) bool {
	switch string(chr) {
	case "a", "e", "i", "u", "o":
		return true
	case "A", "E", "I", "U", "O":
		return true
	default:
		return false
	}
}
