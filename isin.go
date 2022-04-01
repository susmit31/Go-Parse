package main

import "fmt"

var OPS []string = []string{"+", "-", "*", "/"}

func main(){
	x := []int {2,3,4,5,5,4}
	y := []string {"susmit", "honolulu", "moxa", "susmit"}
	type Gurup [T any] []T
	z := Gurup[int] {1,2,3,4}
	fmt.Println(isin(x,500))
	fmt.Println(isin(y,"susmit"))
	fmt.Println(z)
	fmt.Println(indexof(x, 4, 3))
	fmt.Println(y[0][0:3])
}

func isin[T comparable](arr []T, el T) bool{
	for _, c := range arr{
		if c==el{
			return true
		} else {
			return false
		}
	} 
	fmt.Println(OPS[1])
	return false
}

func indexof[T comparable] (arr []T, el T, startat int) int{
	if int(startat) > len(arr)-1{
		return -1
	}
	for i:=startat; i<int(len(arr)); i++{
		if arr[i] == el {
			return i
		}
	}
	return -1
}
