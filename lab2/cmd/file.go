package main

func findKthLargest(array []int, k int) int {
	left := 0
	right := len(array) - 1
	for {
		pivotIndex := partition(array, left, right)
		if pivotIndex == k-1 {
			return nums[pivotIndex]
		} else if pivotIndex > k-1 {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	pivotIndex := left
	for i := left; i < right; i++ {
		if nums[i] > pivot {
			nums[i], nums[pivotIndex] = nums[pivotIndex], nums[i]
			pivotIndex++
		}
	}
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	return pivotIndex
}
