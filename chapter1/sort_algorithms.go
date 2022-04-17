package chapter1

import (
	"fmt"
	// "math"
)

// order 1 asc  2 desc  default is asc
func Insertion(a []int, order int) {
	fmt.Println(a)
	if len(a) <= 0 {
		fmt.Println("数组为空无需排序！")
		return
	}
	for i := 1; i < len(a); i++ {
		//fmt.Println(a[i])
		val := a[i]
		// fmt.Println(val)
		j := i - 1
		for ; j >= 0; j-- {

			left := a[j]
			right := val
			if order == 2 {
				left = val
				right = a[j]
			}

			if left > right {
				a[j+1] = a[j]
				continue
			} else {
				break
			}

		}
		a[j+1] = val

	}

}

// a data array , begin index ,end index
func Merge_Sort(a []int, begin int, end int) {
	fmt.Println(" begin ,end:", begin, end)

	if len(a) <= 0 {
		fmt.Println("数组为空无需排序！")
	}
	if end >= len(a) {
		return
	}
	//begin越界
	if begin < 0 {
		return
	}
	//下标碰头无须比较
	if begin >= end {
		return
	}
	sum := begin + end
	middle := sum / 2

	fmt.Println(" begin ,middle ,end:", begin, middle, end)
	Merge_Sort(a, begin, middle)
	Merge_Sort(a, middle+1, end)
	merge(a, begin, middle, end)

}
func merge(a []int, begin int, middle int, end int) {

	leftLen := middle - begin + 1

	rightLen := end - middle
	arrLeft := make([]int, leftLen+1)
	for i := 0; i < leftLen; i++ {

		arrLeft[i] = a[begin+i]

	}

	arrLeft[leftLen] = 9223372036854775807
	arrRight := make([]int, rightLen+1)
	for j := 0; j < rightLen; j++ {

		arrRight[j] = a[middle+j+1]

	}
	arrRight[rightLen] = 9223372036854775807
	i, j := 0, 0 //i left index   j right index

	for k := begin; k <= end; k++ {
		fmt.Println("------------k:", k)

		if arrLeft[i] <= arrRight[j] {

			a[k] = arrLeft[i]

			i++
			fmt.Println("i:", i, len(arrLeft))
		} else {
			a[k] = arrRight[j]
			j++
			fmt.Println("j:", j, len(arrRight))

		}

	}

}

func Quick_Sort(a []int, begin int, end int) {
	if begin < end {
		mid := partition(a, begin, end)
		fmt.Println("----mid is:", mid)
		fmt.Println("-------mid a:", a)
		Quick_Sort(a, begin, mid-1)
		Quick_Sort(a, mid, end)
	}
}
func partition(a []int, begin int, end int) int {
	x := a[end]
	i := begin - 1
	for k := begin; k < end; k++ {
		if a[k] <= x {
			i = i + 1
			temp := a[i]
			a[i] = a[k]
			a[k] = temp

		}
	}
	tmp := a[i+1]
	a[i+1] = a[end]
	a[end] = tmp
	return i + 1
}
