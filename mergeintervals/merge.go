package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0, len(intervals))
	for i, v := range intervals {
		if i == 0 {
			result = append(result, intervals[0])
		}
		//以前一个进行比较
		if v[0] <= result[len(result)-1][1] { //需要合并
			if v[1] > result[len(result)-1][1]{
				result[len(result)-1][1] = v[1]
			}
		} else { //不需要合并
			result = append(result, v)
		}
	}
	return result
}
