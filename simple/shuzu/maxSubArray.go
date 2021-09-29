package shuzu

//先暴力吧！
func maxSubArray(nums []int) int {
	//找到所有子序列 分别计算和
	var allSub [][]int
	var maxsum = nums[0]
	for i := 0;i < len(nums);i ++ { // i start   j end
		for j := i; j < len(nums); j++ {
			//if SumIJ(nums,i,j) > maxsum {
			//	maxsum = SumIJ(nums,i,j)
			//	continue
			//}
			var s = nums[i:j]
			s = append(s,nums[j])
			allSub = append(allSub,s)
		}
	}
	for _,v := range allSub {
		if SumIJ(v) >= maxsum {
			maxsum = SumIJ(v)
			continue
		}
	}
	return maxsum
}

func SumIJ(nums []int) int {
	var sum int
	for i:= 0;i<len(nums);i++ {
		sum += nums[i]
	}
	return sum
}

func maxSubArray(nums []int) int {
	var max = nums[0]
	for i :=1 ; i < len(nums) ; i ++ {
		if nums[i] + nums[i-1] > nums[i] {   //前者对当前的有帮助 则 “融合”为自身
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

