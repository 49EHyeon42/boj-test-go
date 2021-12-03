package selection

func SelectionSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		minIdx := i
		for j := i; j < len(slice); j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
}
