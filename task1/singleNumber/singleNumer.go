package main

func singleNumber(nums []int) int {
    singleNum := 0
    for _, num := range nums {
        singleNum ^= num
    }
    return singleNum
}