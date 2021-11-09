package shell

func ShellSort(slice []int) {
	for gap := len(slice) / 2; gap > 0; gap /= 2 {
		if gap%2 == 0 {
			gap += 1
		}
		for i := 0; i < gap; i++ {
			for j := i + gap; j <= len(slice)-1; j += gap {
				k, key := 0, slice[j]
				for k = j - gap; i <= k && key < slice[k]; k -= gap {
					slice[k+gap] = slice[k]
				}
				slice[k+gap] = key
			}
		}
	}
}
