package twosum
func twoSum(nums []int, target int) []int {
    newMap := make(map[int]int, len(nums))
    for i, num := range nums {
		tarkey := target - num
		if index, ok := newMap[tarkey]; ok{
			return []int {i, index}
		}
        newMap[num] = i
    }
    return nil

	//暴力查找
/* 	for i, num := range nums{
		complement := target - num
		for j := i + 1; j < len(nums); j++{
			if complement == nums[j] {
				return  []int{i, j}
			}
		} 
	}
	return nil */
}