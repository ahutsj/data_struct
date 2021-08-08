package main

// InsertSort 有些人打扑克时习惯从第二张牌开始，和第一张牌比较，
//第二张牌如果比第一张牌小那么插入到第一张牌前面，这样前两张牌都排好序了，
//接着从第三张牌开始，将它插入到已排好序的前两张牌里，形成三张排好序的牌，
//后面第四张牌继续插入到前面已排好序的三张牌里，直至排序完
//最好情况下，对一个已经排好序的数列进行插入排序，那么需要迭代 N-1 轮，
//并且因为每轮第一次比较，待排序的数就比它左边的数大，那么这一轮就结束了，不需要再比较了，也不需要交换，
//这样时间复杂度为：O(n)
//最坏情况下，每一轮比较，待排序的数都比左边排好序的所有数小，那么需要交换 N-1 次，
//第一轮需要比较和交换一次，第二轮需要比较和交换两次，第三轮要三次，第四轮要四次，这样次数是：1 + 2 + 3 + 4 + ... + N-1，
//时间复杂度和冒泡排序、选择排序一样，都是：O(n^2)
//func InsertSort(n []int) {
//	l := len(n)
//	for i := 1;i < l;i++ {
//		deal := n[i] //要处理的数
//		j := i-1
//		if deal < n[i-1] {
//			for ;j >= 0 && n[j] > deal;j -- {
//				n[j+1] = n[j]
//			}
//			n[j+1] = deal
//		}
//	}
//}
//
//func main() {
//	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
//	InsertSort(list2)
//	fmt.Println(list2)
//}

//数组规模 n 较小的大多数情况下，我们可以使用插入排序，它比冒泡排序，选择排序都快，甚至比任何的排序算法都快
//数列中的有序性越高，插入排序的性能越高，因为待排序数组有序性越高，插入排序比较的次数越少。