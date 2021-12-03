package heap

func HeapSort(slice []int) {
	heapSort2(slice)
}

// * 방법 1
// func heapSort1(slice []int) {
// 	힙 생성
// 	데이터 삽입(insert, push)
// 	데이터 삭제(delete, pop)
// }

// * 방법 2
func heapSort2(slice []int) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		heapify(slice, i, len(slice))
	}
	for i := len(slice) - 1; i >= 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapify(slice, 0, i)
	}
}

func heapify(slice []int, idx, heapSize int) {
	largest, leftIdx, rightIdx := idx, idx*2+1, idx*2+2
	if leftIdx < heapSize && slice[leftIdx] > slice[largest] {
		largest = leftIdx
	}
	if rightIdx < heapSize && slice[rightIdx] > slice[largest] {
		largest = rightIdx
	}
	if largest != idx {
		slice[idx], slice[largest] = slice[largest], slice[idx]
		heapify(slice, largest, heapSize)
	}
}
