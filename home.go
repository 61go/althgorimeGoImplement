package main

import "./alth"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	search := alth.BinarySearch{}
	i := search.Search(arr, 10)
	println("is :")
	println(i)
}
