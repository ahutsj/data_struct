package main

import "fmt"

//选择排序--打扑克牌的时候 排序牌的那种算法，也很垃圾
//每进行一轮迭代，我们都会维持这一轮最小数：
//min 和最小数的下标：minIndex，然后开始扫描，
//如果扫描的数比该数小，那么替换掉最小数和最小数下标，扫描完后判断是否应交换，然后交换
func selectSort(n []int) {
	var l = len(n)
	for i := 0;i < l-1;i ++ {
		var min = n[i]
		var down = i //最小值的下标
		for j := i+1;j < l;j ++ {
			if n[j]<min{
				min = n[j]
				down = j
			}else {
				continue
			}
		}
		n[i],n[down] = n[down],n[i]
	}
}

func main() {
	var a = []int{5, 9, 1, 6, 8, 14, 6, 49,78,-1,0,25, 4, 6, 3}
	selectSort(a)
	fmt.Println(a)
}

//算法的改进
//可以优化算法，使得复杂度减少一半
//每一轮，除了找最小数之外，还找最大数，然后分别和前面和后面的元素交换，这样循环次数减少一半
func selectSortBetter(n []int) {
	var l = len(n)
	for i := 0;i<l;i++ {

	}
}