package leetcode

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)

	var isOdd = (n+m)%2 != 0
	fmt.Println("isOdd", isOdd)

	if n == 0 {
		nums1 = nums2
		n = m
	}
	if m == 0 {
		if isOdd {
			return float64(nums1[(n-1)/2])
		}
		return float64((nums1[n/2-1] + nums1[n/2+1]) / 2)
	}

	var c1, c2 int
	if nums1[0] < nums2[0] {
		c1, c2 = nums1[0], nums2[0]
	}

	fmt.Printf("c1 %d, c2 %d\n", c1, c2)

	var curRank = 0
	for i, j := 1, 1; i < n || j < m; {
		if isOdd && curRank == (n+m-1)/2 {
			fmt.Printf("i=%d, j=%d\n", i, j)
			break
		} else if !isOdd && curRank == (n+m)/2 {
			fmt.Printf("i=%d, j=%d\n", i, j)
			break
		}

		if i == n {
			if isOdd {
				c1 = nums2[(n+m-1)/2-i]
				c2 = c1
				curRank++
			} else {
				c1, c2 = nums2[(n+m)/2-1-i], nums2[(n+m)/2+1-i]
				curRank++
			}
		} else {
			if c1 < nums1[i] {
				c1 = c2
				c2 = nums1[i]
				curRank++
			} else if c1 < nums1[i] && c2 > nums1[i] {
				c1 = nums1[i]
				curRank++
			}
			i++
		}

		if j == m {
			if isOdd {
				c1 = nums1[(n+m-1)/2-j]
				c2 = c1
				curRank++
			} else {
				c1, c2 = nums1[(n+m)/2-1-j], nums1[(n+m)/2+1-j]
				curRank++
			}
		} else {
			if c2 < nums2[j] {
				c1 = c2
				c2 = nums2[j]
				curRank++
			} else if c1 < nums2[j] && c2 > nums2[j] {
				c1 = nums2[j]
				curRank++
			}

			j++
		}
	}

	fmt.Println("curRank", curRank)
	fmt.Printf("c1 %d, c2 %d\n", c1, c2)
	return float64((c1 + c2) / 2)
}

func testAsk() {
	fmt.Println("testing")
}

// divide by two
func mergeSort(nums []int, l, r int) {
	if l < r {
		m := l + (r-l)/2

		// sort left part
		mergeSort(nums, l, m)
		// sort right part
		mergeSort(nums, m+1, r)

		// merge left and right parts
		ll := m - l + 1       // left array len
		rl := r - (m + 1) + 1 // right array len

		var L = make([]int, ll)
		var R = make([]int, rl)

		for i := 0; i < ll; i++ {
			L[i] = nums[l+i]
		}

		for i := 0; i < rl; i++ {
			R[i] = nums[m+1+i]
		}

		// merge by comparing left and right parts
		var a, b, c int
		c = l // begining of the current mergesort session
		for a < ll && b < rl {
			if L[a] < R[b] {
				nums[c] = L[a]
				c++
				a++
			} else {
				nums[c] = R[b]
				c++
				b++
			}
		}

		// copy rest of left
		for a < ll {
			nums[c] = L[a]
			a++
			c++
		}

		for b < rl {
			nums[c] = R[b]
			b++
			c++
		}
	}
}

func partition(nums []int, low, high int) (i int) {
	piv := nums[high]

	i = low - 1
	for j := low; j < high; j++ {
		if nums[j] < piv {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}

	i++
	nums[i], nums[high] = nums[high], nums[i]

	return i
}

// recursive version
func quickSort(nums []int, low, high int) []int {
	if low < high {
		i := partition(nums, low, high)

		quickSort(nums, low, i-1)
		quickSort(nums, i+1, high)
	}

	return nums
}
