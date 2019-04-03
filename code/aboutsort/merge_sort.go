package aboutsort

// 归并排序
func MergeSort(sli []int) {
	length := len(sli)
	if length <= 1 {
		return
	}

	mergeSort(sli, 0, length-1)
}

func mergeSort(sli []int, start, end int)  {
	if start >= end {
		return
	}

	mid := (start+end)/2

	mergeSort(sli, start, mid)
	mergeSort(sli, mid+1, end)
	merge(sli, start, mid, end)
}

func merge(sli []int, start, mid, end int)  {
	tmpSli := make([]int, end-start+1)

	i := start
	j := mid+1
	k := 0
	for ; i <= mid && j <= end; k++{
		if sli[i] < sli[j] {
			tmpSli[k] = sli[i]
			i++
		}else {
			tmpSli[k] = sli[j]
			j++
		}
	}

	for ; i <= mid; i++{
		tmpSli[k] = sli[i]
		k++
	}

	for ; j <= end; j++{
		tmpSli[k] = sli[j]
		k++
	}
	copy(sli[start:end+1], tmpSli)
}


