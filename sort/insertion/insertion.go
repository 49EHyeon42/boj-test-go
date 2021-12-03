package insertion

func InsertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
		}
	}
}
