package plusone

func plusOne(digits []int) []int {
	//从后往前遍历,完成加一操作
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i]++
		}

		if digits[i] == 10 && i-1 >= 0 {
			digits[i] = 0
			digits[i-1]++
		}
	}
	if digits[0] == 10 {
		//扩容
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}
