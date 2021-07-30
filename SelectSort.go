package main

//选择排序--打扑克牌的时候 排序牌的那种算法，也很垃圾
//每进行一轮迭代，我们都会维持这一轮最小数：
//min 和最小数的下标：minIndex，然后开始扫描，
//如果扫描的数比该数小，那么替换掉最小数和最小数下标，扫描完后判断是否应交换，然后交换
//func selectSort(n []int) {
//	var l = len(n)
//	for i := 0;i < l-1;i ++ {
//		var min = n[i]
//		var down = i //最小值的下标
//		for j := i+1;j < l;j ++ {
//			if n[j]<min{
//				min = n[j]
//				down = j
//			}else {
//				continue
//			}
//		}
//		n[i],n[down] = n[down],n[i]
//	}
//}
//
//func main() {
//	var a = []int{5, 9, 1, 6, 8, 14, 6, 49,78,-1,0,25, 4, 6, 3}
//	selectSort(a)
//	fmt.Println(a)
//}
//
//// SelectGoodSort 算法的改进
////可以优化算法，使得复杂度减少一半
////每一轮，除了找最小数之外，还找最大数，然后分别和前面和后面的元素交换，这样循环次数减少一半
//func SelectGoodSort(list []int) {
//	n := len(list)
//	// 只需循环一半
//	for i := 0; i < n/2; i++ {
//		minIndex := i // 最小值下标
//		maxIndex := i // 最大值下标
//		// 在这一轮迭代中要找到最大值和最小值的下标
//		for j := i + 1; j < n-i; j++ {
//			// 找到最大值下标
//			if list[j] > list[maxIndex] {
//				maxIndex = j // 这一轮这个是大的，直接 continue
//				continue
//			}
//			// 找到最小值下标
//			if list[j] < list[minIndex] {
//				minIndex = j
//			}
//		}
//		if maxIndex == i && minIndex != n-i-1 {
//			// 如果最大值是开头的元素，而最小值不是最尾的元素
//			// 先将最大值和最尾的元素交换
//			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
//			// 然后最小的元素放在最开头
//			list[i], list[minIndex] = list[minIndex], list[i]
//		} else if maxIndex == i && minIndex == n-i-1 {
//			// 如果最大值在开头，最小值在结尾，直接交换
//			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
//		} else {
//			// 否则先将最小值放在开头，再将最大值放在结尾
//			list[i], list[minIndex] = list[minIndex], list[i]
//			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
//		}
//	}
//}