package practice

import "fmt"

func NextPermutation(nums []int) {
	if isMaxPermutation(nums) {
		sort(nums)
		return
	}

	var subslice []int
	var newPosition int

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			subslice = nums[i:]
			newPosition = i - 1
			break
		}
	}

	fmt.Println("Swap Position", newPosition)

	doSwap(nums, newPosition)

	fmt.Println("Subslice after swap", subslice)

	sort(subslice)

	fmt.Println("Subslice after sort", subslice)

}

func doSwap(nums []int, pos int) {

	switchFromPos := pos

	for i := pos + 1; i < len(nums); i++ {
		if nums[i] > nums[pos] {
			switchFromPos = i
		}
	}

	if switchFromPos != pos {
		temp := nums[switchFromPos]
		nums[switchFromPos] = nums[pos]
		nums[pos] = temp
	}
}

func isMaxPermutation(nums []int) bool {

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			return false
		}
	}

	return true
}

func sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
}
