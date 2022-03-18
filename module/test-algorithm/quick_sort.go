package main

import "fmt"

/**
快速排序
 */
// 8, 10, 3, 5, 2, 7, 4, 9, 1, 6
// 6, 10, 3, 5, 2, 7, 4, 9, 1, 8
// 6, 8, 3, 5, 2, 7, 4, 9, 1, 10
func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	left, right := low, high
	pivot := arr[left] // 基数， 取第一个

	// 从左到右
	for left < right {
		// 从右边开始
		for left < right && arr[right] >= pivot {
			right --
		}
		// 交换
		if left < right {
			arr[left] = arr[right]
		}
		//arr[right] = pivot
		// 从左开始
		for left < right && arr[left] < pivot {
			left ++
		}
		// 交换
		if left < right {
			arr[right] = arr[left]
		}
		if left <= right {
			arr[left] = pivot
		}
	}

	// 继续排序左边
	quickSort(arr, low, right - 1)
	// 继续排序右边
	quickSort(arr, right + 1, high)
}
func main() {
	value := []int{1,2,3,4,5,6,7,8,9}
	quickSort(value, 0, len(value) - 1)
	fmt.Println("0 ..... 10 add total ", value)
}
