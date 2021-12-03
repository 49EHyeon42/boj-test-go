package merge

func MergeSort(slice []int) []int {
	return mergeSort(slice)
}

func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	first := mergeSort(slice[:len(slice)/2])
	second := mergeSort(slice[len(slice)/2:])
	return merge(first, second)
}

func merge(first, second []int) []int {
	sorted, i, j := []int{}, 0, 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			sorted, i = append(sorted, first[i]), i+1
		} else {
			sorted, j = append(sorted, second[j]), j+1
		}
	}
	for ; i < len(first); i++ {
		sorted = append(sorted, first[i])
	}
	for ; j < len(second); j++ {
		sorted = append(sorted, second[j])
	}
	return sorted
}
