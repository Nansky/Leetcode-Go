// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

// Constraints:

// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return float64(0)
	}

	var med float64
	var newArr []int

	newArr = append(nums1, nums2...)

	sortedArr := sortArr(newArr)
	fmt.Println(sortedArr)

	mid := (len(sortedArr) - 1) / 2

	if len(sortedArr)%2 == 0 {
		med = float64(sortedArr[mid]+sortedArr[mid+1]) / 2
	} else {
		med = float64(sortedArr[mid])
	}

	return med
}

func sortArr(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	left := sortArr(nums[:mid])
	right := sortArr(nums[mid:])

	res := join(left, right)

	return res
}

func join(leftArr, rightArr []int) []int {
	m := make([]int, 0)
	i, j := 0, 0

	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] < rightArr[j] {
			m = append(m, leftArr[i])
			i++
		} else {
			m = append(m, rightArr[j])
			j++
		}
	}

	for ; i < len(leftArr); i++ {
		m = append(m, leftArr[i])
	}

	for ; j < len(rightArr); j++ {
		m = append(m, rightArr[j])
	}

	return m
}