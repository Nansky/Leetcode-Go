// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

func maxArea(height []int) int {
	var maxArea int
	i := 0
	j := len(height) - 1

	for i < j {
		distance := j - i
		minBar := minHeight(height[i], height[j])

		area := minBar * distance

		if maxArea < area {
			maxArea = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}

	return maxArea
}

func minHeight(b1, b2 int) int {
	if b1 < b2 {
		return b1
	}

	return b2
}