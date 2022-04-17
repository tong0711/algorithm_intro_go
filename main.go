// lk project main.go
package main

import (
	"fmt"
	"math"
	"math/cmplx"

	// "math/rand"
	// "time"

	"example.com/algorithm/chapter1"
	"example.com/algorithm/chapter4"
)

func add(x int, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//var c, python, java bool
var c, python, java = true, false, "no!"

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

type Vertex struct {
	X int
	Y int
}

func main() {
	// fmt.Println("Hello World!")
	// fmt.Println("The time is", time.Now())
	// fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Println("My cos60 number is", math.Cos(60))
	// fmt.Println("50 +20  number is", add(50, 20))
	// //fmt.Println("chapter1:50 +20  number is", chapter1.Add(50, 20))
	// a, b := swap("a", "b")
	// fmt.Println("swap  a and b   is:", a, b)
	// //fmt.Println("split 17 is：", split(17))
	// fmt.Println(split(17))
	// //var i int
	// var i int = 2
	// fmt.Println(i, c, python, java)
	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)
	// var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var k uint = uint(f)
	// k1 := 14
	// fmt.Println(x, y, k, k1)
	// v := 42 // change me!
	// fmt.Printf("v is of type %T\n", v)
	// const World = "世界"
	// fmt.Println("Hello", World)
	// fmt.Println("Happy", Pi, "Day")

	// const Truth = true
	// fmt.Println("Go rules?", Truth)
	// fmt.Println(needInt(Small))
	// fmt.Println(needFloat(Small))
	// fmt.Println(needFloat(Big))
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)
	// sum = 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)
	// fmt.Println(sqrt(2), sqrt(-4))
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )
	// j := 42
	// p := &i         // point to i
	// fmt.Println(*p) // read i through the pointer
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i

	// p = &j          // point to j
	// *p = *p / 37    // divide j through the pointer
	// fmt.Println(j)  // see the new value of jp := &i         // point to i
	// fmt.Println(*p) // read i through the pointer
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i

	// p = &j         // point to j
	// *p = *p / 37   // divide j through the pointer
	// fmt.Println(j) // see the new value of j
	// fmt.Println(Vertex{1, 2})
	var as = [12]int{5, 24, 3, 44, 51, 6, 71, 81, 29, 110, 11, 23}
	slice := as[:]
	// chapter1.Insertion(slice, 2)
	chapter1.Merge_Sort(slice, 0, 11)
	var as1 = [12]int{5, 24, -3, 44, 51, -6, 71, -81, 29, -110, 11, 23}
	slice1 := as1[:]
	fmt.Println("slice1:", slice1)
	result := chapter4.Max_subarray(slice1, 0, 11)

	//var asp *[3]int = &as

	fmt.Println(slice1)
	fmt.Println(" max subarray:", result)
	var as2 = [12]int{5, 24, -3, 44, 51, -6, 71, -81, 29, -110, 11, 23}
	slice2 := as2[:]
	chapter1.Quick_Sort(slice2, 0, len(slice2)-1)
	fmt.Print(" slice quick sort:", slice2)

}
