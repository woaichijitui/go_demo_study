package slice

import "fmt"

// K1 扩容影响机制
func K1() {

	intA := make([]int, 3, 4)
	intA[0], intA[1] = 1, 2
	// B 和 A 指向同一个地址
	intB := intA[:]
	fmt.Printf("A 扩容前 ：指针：%p ， 长度： %d ，容量： %d  ；B 扩容前 ：%p ， 长度： %d ，容量： %d  \n ", intA, len(intA), cap(intA), intB, len(intB), cap(intB))

	//	扩容：
	intA = append(intA, 3, 4)
	fmt.Printf("A 扩容后 ：指针：%p ， 长度： %d ，容量： %d  ；B 扩容后 ：%p ， 长度： %d ，容量： %d  \n ", intA, len(intA), cap(intA), intB, len(intB), cap(intB))

}
