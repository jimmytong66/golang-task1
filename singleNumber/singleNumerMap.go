
package main

func singleNumberMap(nums []int) int {
    singleNum := 0
    countMap := make(map[int]int, len(nums)/2+1)
    for _, num := range nums {
        countMap[num]++
    }
    for k, v := range countMap{
        if v == 1 {
            singleNum = k
        }
    }
    return singleNum
}