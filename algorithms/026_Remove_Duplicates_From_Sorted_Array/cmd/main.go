package main

func RemoveDuplicates() int {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	k := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			arr[k] = arr[i]
			k = k + 1
		}
	}

	return k

}
