package removeduplicate

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			nums[slow+1] = nums[fast]
			slow++
		}
	}
	return slow + 1
}
