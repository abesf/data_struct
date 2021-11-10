package test

import "testing"

func TestBinarySearch(t *testing.T) {

}
func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}
func search2(nums []int, target int) int {
	return -1
}
