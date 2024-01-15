package src

func BinarySearchMatrix(nums []int, target int) int {
	low, high := 0, len(nums)-1
	if low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
