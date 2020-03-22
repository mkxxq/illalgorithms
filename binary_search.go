package illalgorithms

func BinarySearch(array []int, target int) int {
	res := -1
	if len(array) < 1 {
		return res
	}
	low := 0
	high := len(array) - 1

	for array[low] < array[high] {
		mid := (low + high) / 2
		if array[mid] > target {
			high = mid
		} else if array[mid] < target {
			low = mid
		} else {
			res = mid
			break
		}
	}
	return res
}
