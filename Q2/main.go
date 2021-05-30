package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BubbleSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func CariMedian(arr []int) (cari int) {
	tot := len(arr)
	if tot%2 == 0 {
		bagi := tot / 2
		cari = (arr[(bagi)-1] + arr[bagi]) / 2
		// fmt.Println(cari)
	} else {
		cari = arr[(tot)/2]
		// fmt.Println(cari)
	}
	// fmt.Print(cari)
	return
}

func FindBlock(totblock, start int, arr []int) (hasil []int) {
	listarr := arr[start:len(arr)]
	for i, row := range listarr {
		if i == totblock {
			break
		} else {
			hasil = append(hasil, row)
		}
	}
	// fmt.Print(hasil)
	return
}

func MovingMedian(arr []int) {
	// arr_jadi := []int{}
	acuan := arr[0]
	listnumber := arr[1:len(arr)]
	listmedian := []int{}
	x := 1
	y := 0
	for i, row := range listnumber {
		if i == 0 {
			listmedian = append(listmedian, row)
		} else if x >= acuan {
			if (acuan + (y - 1)) == len(listnumber) {
				//terakhir
				break
			} else {
				cariblock := FindBlock(acuan, y, listnumber)
				sortbuble := BubbleSort(cariblock)
				carimedian := CariMedian(sortbuble)
				listmedian = append(listmedian, carimedian)
			}
			y++
		} else if i > 0 {
			cariblock := FindBlock(x, 0, listnumber)
			sortbuble := BubbleSort(cariblock)
			carimedian := CariMedian(sortbuble)
			listmedian = append(listmedian, carimedian)
		}
		x++
	}

	//CONVERT TO STRING
	var HASILMEDIAN []string
	for _, i := range listmedian {
		HASILMEDIAN = append(HASILMEDIAN, strconv.Itoa(i))
	}
	fmt.Println(strings.Join(HASILMEDIAN, ", "))
}
func main() {
	input := []int{5, 2, 4, 6}
	MovingMedian(input)
}
