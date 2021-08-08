package main

//import (
//	"fmt"
//)
//
//func ShellSort(n []int) {
//	l := len(n)
//
//	for k := l/2;k >= 1;k /= 2 { //步长
//		//插入排序
//		for i := l-k;i>=0;i -= k {
//			 //插入排序中待处理的数字
//			j := i-k
//			for ;j >= 0;j -= k {
//				if n[j] > n[i] {
//					n[i],n[j] = n[j],n[i]
//					continue
//				}
//			}
//		}
//
//	}
//}
//
//func main () {
//	list  := []int{2,4,1,0,9,8,0,1}
//	ShellSort(list)
//	fmt.Println(list)
//}
