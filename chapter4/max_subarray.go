package chapter4

import "fmt"

// "math"

type result struct {
	low  int
	high int
	sum  int
}

func Max_subarray(a []int, low int, high int) result {
	/* /fmt.Print("max subarray:", a*/
	fmt.Println("-------------subarray:", low, high)
	if low == high {
		return result{low, high, a[low]}
	}
	mid := (low + high) / 2
	ls := Max_subarray(a, low, mid)
	rs := Max_subarray(a, mid+1, high)
	cs := findCrossMaxSunArray(a, low, mid, high)
	if ls.sum >= rs.sum && ls.sum >= cs.sum {
		return ls
	}
	if rs.sum >= ls.sum && rs.sum >= cs.sum {
		return rs
	}
	return cs
}
func findCrossMaxSunArray(a []int, low int, mid int, high int) result {
	fmt.Println(" cross low,mid,high:", low, mid, high)
	lsum := -9223372036854775807
	sum := 0
	ml := mid

	for i := mid; i >= low; i-- {
		index := i
		if index < 0 {
			index = 0
		}
		sum = sum + a[index]
		if sum > lsum {
			lsum = sum
			ml = index
		}
	}

	rsum := -9223372036854775807
	sum = 0
	mr := mid

	for j := mid + 1; j <= high; j++ {
		index := j
		if index > high {
			index = high
		}
		sum = sum + a[index]
		if sum > rsum {
			rsum = sum
			mr = index
		}
	}
	return result{ml, mr, lsum + rsum}
}
