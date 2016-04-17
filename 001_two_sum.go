package leetcodeingo

func twoSum(nums []int, target int) []int {

	var ret []int
	dict := make(map[int]int)
	for i, c := range nums {
		pre, ok := dict[target-c]
		if ok {
			ret = append(ret, pre)
			ret = append(ret, i)
			return ret
		}
		dict[c] = i
	}
	return ret
}
