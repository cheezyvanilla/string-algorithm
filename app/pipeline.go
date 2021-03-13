package app

import (
	"fmt"
	"time"
)

func main(){
	arr := []int{3,7,5,2}

	sorted := sort(arr)
	fmt.Println(sorted)
}

func sort(n []int)[]int{
	if len(n) == 1{
		return n
	}

	arr1 := n[0:(len(n)/2)]
	arr2 := n[(len(n)/2):len(n)]

	arr1 = sort(arr1)
	arr2 = sort(arr2)

	return merge(arr1, arr2)

}

func merge(n1, n2 []int) []int{
	totalLength := len(n1) + len(n2)
	n := make([]int, totalLength)

	i := 0
	for len(n1) > 0 && len(n2) > 0 {
		if n1[0] > n2[0]{
			n[i] = n2[0]
			n2 = n2[1:]
		}else{
			n[i] = n1[0]
			n1 = n1[1:]
		}
		fmt.Println(n, n1, n2)
		i++
	}

	for len(n1) > 0 {
		n[i] = n1[0]
		n1 = n1[1:]
	}

	for len(n2) > 0 {
		n[i] = n2[0]
		n2 = n2[1:]
	}

	return n
}

type test struct{
	name string
	mulai time.Time
}