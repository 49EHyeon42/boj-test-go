package quick

func QuickSort(slice []int) {
	quickSort(slice, 0, len(slice)-1)
}

func quickSort(slice []int, left, right int) {
	if left < right {
		i, j, pivot := left, right, slice[(left+right)/2]
		for i < j {
			for slice[j] > pivot {
				j -= 1
			}
			for i < j && slice[i] < pivot {
				i += 1
			}
			slice[i], slice[j] = slice[j], slice[i]
		}
		quickSort(slice, left, i-1)
		quickSort(slice, i+1, right)
	}
}
