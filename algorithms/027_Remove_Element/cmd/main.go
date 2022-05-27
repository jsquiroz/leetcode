package main

func RemoveDuplicates(arr []int, val int) int {

	k := 0

	for i := 0; i < len(arr); i++ {
		if val != arr[i] {
			arr[k] = arr[i]
			k = k + 1
		}
	}

	return k
}
