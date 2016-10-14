// Package main provides ...
package newDir

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var randFeed *rand.Rand
var ints []int

const count = 20

func init() {
	randFeed = rand.New(rand.NewSource(time.Now().UnixNano()))
	ints = randFeed.Perm(count)
}

func JustDemo() {
	println("//<<-------------------------JustDemo start-----------")
	start := time.Now()
	a := 5 >> 1
	fmt.Printf("a: %+v\n", a)
	// insertionSortDEmo()
	bubbleSortDemo()
	fmt.Printf("JustDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------JustDemo end----------->>")
}

func insertionSortDEmo() {
	println("//<<-------------------------insertionSortDEmo start-----------")
	start := time.Now()
	ints := randFeed.Perm(count)
	fmt.Printf("ints: %+v\n", ints)

	fmt.Printf("insertionSortDEmo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------insertionSortDEmo end----------->>")
}

func bubbleSortDemo() {
	println("//<<-------------------------bubbleSortDemo start-----------")
	start := time.Now()
	fmt.Printf("o ints: %+v\n", ints)
	for i := len(ints) - 1; i > 1; i-- {
		swap := false
		for j := 0; j < i; j++ {
			if ints[j+1] < ints[j] {
				ints[j+1], ints[j] = ints[j], ints[j+1]
				swap = true
			}
		}

		//all elements are in order
		if swap == false {
			break
		}

		// for j := i - 1; j > 0; j-- {
		// 	println(j)
		// 	if ints[j+1] > ints[j] {
		// 		ints[j+1], ints[j] = ints[j], ints[j+1]
		// 	}
		// }
	}
	fmt.Printf("ints: %+v\n", ints)

	fmt.Printf("bubbleSortDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------bubbleSortDemo end----------->>")
}

func quickSortDemo() {
	arr := make([]int, count)
	copy(arr, ints)
	fmt.Printf("o arrs: %+v\n", arr)

	var recurse func(left int, right int)
	var partition func(left int, right int, pivot int) int

	partition = func(left int, right int, pivot int) int {
		v := arr[pivot]
		right--
		arr[pivot], arr[right] = arr[right], arr[pivot]

		for i := left; i < right; i++ {
			if arr[i] <= v {
				arr[i], arr[left] = arr[left], arr[i]
				left++
			}
		}

		arr[left], arr[right] = arr[right], arr[left]
		return left
	}

	recurse = func(left int, right int) {
		if left < right {
			pivot := (right + left) / 2
			pivot = partition(left, right, pivot)
			recurse(left, pivot)
			recurse(pivot+1, right)
		}
	}

	println("//<<-------------------------quickSortDemo start-----------")
	start := time.Now()
	recurse(0, len(arr)-1)

	fmt.Printf("arrs: %+v\n", arr)

	fmt.Printf("quickSortDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------quickSortDemo end----------->>")
}

// func getTokenDemo() {
// 	println("//<<-------------------------getTokenDemo start-----------")
// 	start := time.Now()
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	// a := r.Intn(1000)
//
// 	fmt.Printf("a: %+v\n", a)
// 	fmt.Printf("getTokenDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
// 	println("//---------------------------getTokenDemo end----------->>")
// }
func SwapDemo(a int, b int) {
	println("//<<-------------------------SwapDemo start-----------")
	start := time.Now()
	fmt.Printf("a: %b\n", a)
	fmt.Printf("b: %b\n", b)
	a = (a ^ b)
	b = (a ^ b)
	a = (a ^ b)
	println("after")
	fmt.Printf("a: %b\n", a)
	fmt.Printf("b: %b\n", b)
	fmt.Printf("SwapDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------SwapDemo end----------->>")
}

func ZeroAtEndDemo() {
	println("//<<-------------------------ZeroAtEndDemo start-----------")
	start := time.Now()
	// arr := []int{0, 1, 2, 3, 0, 8}
	arr := []int{0, 0, 1, 2, 3, 0, 8, 0}
	arrLen := len(arr)
	j := 0
	log.Println("before: works %v", arr)

	for i := 0; i < arrLen; i++ {
		if arr[i] == 0 {
			j = i + 1
			for j < arrLen {
				if arr[j] != 0 {
					arr[i], arr[j] = arr[j], arr[i]
					j = arrLen //just skip next swap for this arr[i]
				}
				j++
			}
		}
	}

	// p := 0
	// for p < arrLen {
	// 	if arr[p] == 0 {
	// 		break
	// 	}
	// 	p++
	// }
	// q := p
	// for q < arrLen {
	// 	if arr[q] != 0 {
	// 		arr[p] = arr[q]
	// 		p++
	// 	}
	// 	q++
	// }
	// for p < arrLen-1 {
	// 	arr[p] = 0
	// 	p++
	// }
	// for (int i = 0; i <nums.size(); ++i)
	// {
	// 	if (nums[i]==0)
	// 	{
	// 		j=i+1;
	// 		while(j<nums.size()){
	// 			if (nums[j]!=0)
	// 			{
	// 				nums[j]=nums[i]^nums[j];
	// 				nums[i]=nums[i]^nums[j];
	// 				nums[j]=nums[i]^nums[j];
	// 				j=nums.size();
	// 			}
	// 			j++;
	// 		}
	// 	}
	// }
	log.Println("after : works %v", arr)

	fmt.Printf("ZeroAtEndDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ZeroAtEndDemo end----------->>")
}

func removeInstanceOfIndexingDemo() {
	println("//<<-------------------------removeInstanceOfIndexingDemo start-----------")
	start := time.Now()
	// nums := []int{3, 2, 2, 3, 4, 5}
	nums := []int{3, 2, 2, 3}
	input := 2
	numsLen := len(nums)
	target := nums[input]
	tail := numsLen - 1
	for i := 0; i < numsLen; i++ {
		if nums[i] == target {
			for tail >= i {
				if nums[tail] == target {
					tail--
					numsLen--
					continue
				}
				nums[i], nums[tail] = nums[tail], nums[i]
				tail--
				numsLen--
				break
			}

		}
		if i > tail {
			break
		}
	}

	log.Printf("numsLen: %+v\n", numsLen)
	log.Printf("numsLen: %+v\n", nums[:numsLen])

	fmt.Printf("removeInstanceOfIndexingDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------removeInstanceOfIndexingDemo end----------->>")
}
