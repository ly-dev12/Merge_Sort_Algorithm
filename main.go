package main

import "fmt"

func MergeSort(array []int, m int) []int {
	//fmt.Println(array)

	if m == 1 {
		//fmt.Println(array)
		return array
	}
	op := m / 2
	x1 := make([]int, 0)
	x2 := make([]int, 0)

	x1 = append(x1, array[0:op]...)
	x2 = append(x2, array[op:]...)
	//MergeSort(x1, op)
	//MergeSort(x2, op)
	return MergeIU(MergeSort(x1, op), MergeSort(x2, op))
}

func MergeIU(a []int, b []int) []int {
	xr := make([]int, 0)
	xr = append(xr, a...)
	xr = append(xr, b...)
	//fmt.Println(xr)

	for vi := 0; vi < len(xr); vi++ {
		xd := xr[vi]
		fmt.Println(xr, xd)
		for i := 0; i < len(xr); i++ {
			/*if xd == xr[i] {
				break
			}*/
			if xd < xr[i] {
				xr[vi] = xr[i]
				xr[i] = xd
				vi--
				break
			}
		}
	}
	fmt.Println("XR: ", xr)
	return xr
}

func main() {
	array := []int{5, 2, 6, 7, 3, 1, 8, 4, 15, 9, 14, 20, 17, 18}
	MergeSort(array, len(array))
}
