package bubble

func BubbleSort1(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func BubbleSort2(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		swap := false
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				swap, slice[j], slice[j+1] = true, slice[j+1], slice[j]
			}
		}
		if !swap {
			break
		}
	}
}
