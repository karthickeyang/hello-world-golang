/*

this is rogram is mainly used for average amrk
hi i am using 
*/

package main

import (
	"fmt"
)

func main() {

	var marks = []int {98, 89, 67, 99, 100}
	var size int
	size = 5
	var for_pointer int 
	for_pointer = 10
	fmt.Println("hi i am so happyt",&for_pointer)
	ave := Average(marks, size)
	fmt.Println("average of your 10th marks %d", ave)

}

func Average(ave []int, count int) float32 {
	var sum int
	var ave1 float32
	for i := 0; i < count; i++ {

		sum += ave[i]
	}

	ave1 = float32(sum / count)
	return ave1
}
