package shuzu

func containsDuplicate(nums []int) bool {
	exist := make(map[int]int)
	for i := 0;i < len(nums);i ++{
		_,ok := exist[nums[i]]
		if !ok {
			exist[nums[i]] = 1
			continue
		} else {
			return true
		}
	}
	return false
}
