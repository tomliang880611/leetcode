package leetcode

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	low := 0
	high := len(nums) - 1

	for high-low > 1 {
		piv := (low + high) / 2

		if nums[piv] == target {
			return piv
		}

		if nums[low] == target {
			return low
		}
		if nums[high] == target {
			return high
		}

		if nums[piv] > nums[high] {
			if nums[low] < target && target < nums[piv] {
				high = piv
			} else {
				low = piv
			}
		} else if nums[low] > nums[piv] {
			if nums[piv] < target && target < nums[high] {
				low = piv
			} else {
				high = piv
			}
		} else { // low -> high
			if nums[piv] > target {
				high = piv
			} else {
				low = piv
			}
		}
	}

	if nums[low] == target {
		return low
	}
	if nums[high] == target {
		return high
	}
	return -1
}
