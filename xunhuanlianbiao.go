package main
//
//import "fmt"
//
//type Ring struct {
//	prev,next   *Ring
//	value       interface{}
//}
//
//func NewRing(n int) *Ring {
//	r := new(Ring)
//	r.value = 1
//	p := r
//	for i := 1;i < n ;i ++ {
//		p.next = &Ring{prev: p,value: i+1}
//		p = p.next
//	}
//	p.next = r
//	r.prev = p
//	return r
//}
//
//func main() {
//	r := NewRing(5)
//	fmt.Println(r.value,r.next.value,r.next.next.value,r.prev.value,r.prev.next.value)
//}