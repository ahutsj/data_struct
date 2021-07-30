package main

//冒泡排序是效率较低的排序算法，可以说是最慢的排序算法
//func bubbleSort(n []int){
//	var l = len(n)
//	didSwap := false
//	for j := l;j > 0; j -- {
//		for i := 1;i <j ; i++ {
//			if n[i-1] > n[i] {
//				n[i-1],n[i] = n[i],n[i-1]
//				didSwap = true
//			}else {
//				continue
//			}
//		}
//
//		//如果在一轮中没有交换过，那么已经排好序了，直接返回
//		//也就是说在最好的情况下：对已经排好序的数列进行冒泡排序，只需比较 N 次
//		//最好时间复杂度从 O(n^2) 骤减为 O(n)
//		if !didSwap {
//			return
//		}
//	}
//}
//
//func main() {
//	var a = []int{5, 9, 1, 6, 8, 14, 6, 49,78,-1,0,25, 4, 6, 3}
//	bubbleSort(a)
//	//因为切片会原地排序，排序函数不需要返回任何值，处理完后可以直接打印
//	fmt.Println(a)
//}
