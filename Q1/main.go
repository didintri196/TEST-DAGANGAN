package main

import "fmt"

func SumAll(arr []int) (tot int) {
	for _, a := range arr {
		tot += a
	}
	return
}
func SumMultiplier(arr []int) bool {
	sumall := SumAll(arr)
	fmt.Println(sumall)
	terbesar1, terbesar2 := SearchMax(arr)
	kali2 := sumall * 2
	total := terbesar1 * terbesar2
	if kali2 > total {
		return false
	} else {
		return true
	}
}

func SearchMax(arr []int) (terbesar1, terbesar2 int) {

	for i, row := range arr {
		if terbesar1 < row {
			terbesar2 = terbesar1
			terbesar1 = row
		}
		if i == len(arr) {
			break
		}
	}
	return
}

func main() {
	data := []int{1, 1, 2, 10, 3, 1, 12}
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(SumMultiplier(data))
}
