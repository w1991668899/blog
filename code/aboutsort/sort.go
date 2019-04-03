package aboutsort

// 冒泡排序 从小到大
func BubbleSort1(sli []int)[]int {
	var length = len(sli)
	if length <= 1 {
		return sli
	}

	for i := 0; i < length -1; i++ {
		// 提前退出标志
		flag := false
		for j := 0; j < length-i-1; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
				//此次冒泡有数据交换
				flag = true
			}
		}
		// 如果没有交换数据，提前退出
		if !flag {
			break
		}
	}

	return sli
}

// 冒泡排序 从大到小
func BubbleSort2(sli []int) []int {
	length := len(sli)
	if length <= 1 {
		return sli
	}

	for i := 0; i < length; i++{
		var flag bool
		for j := 0; j < length-i-1; j++{
			if sli[j] < sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	return sli
}

// 插入排序 从小到大
func InsertSort1(sli []int)[]int{

	var length = len(sli)
	if length <= 1 {
		return sli
	}

	for i := 1; i < length; i++{
		value := sli[i]
		j := i - 1
		for ; j >= 0; j--{
			if sli[j] > value{
				sli[j+1] = sli[j]
			}else {
				break
			}
		}

		sli[j+1] = value
	}
	return sli
}

// 插入排序 大到小
func InsertSort2(sli []int) []int {
	length := len(sli)
	if length <= 1 {
		return sli
	}

	for i := 1; i < length; i++{
		value := sli[i]
		j := i - 1
		for ; j >= 0; j--{
			if sli[j] < value {
				sli[j+1] = sli[j]
			}else {
				break
			}
		}
		sli[j+1] = value
	}
	return sli
}

// 选择排序 从小到大
func SelectSort1(sli []int)[]int{
	length := len(sli)
	if length <= 1 {
		return sli
	}

	for i := 0; i < length; i++{
		minIndex := i
		for j := i + 1; j < length; j++{
			if sli[j] < sli[minIndex] {
				minIndex = j
			}
		}

		sli[i], sli[minIndex] = sli[minIndex], sli[i]
	}

	return sli
}

// 选择排序 从大到小
func SelectSort2(sli []int)[]int  {
	length := len(sli)
	if length <= 1 {
		return sli
	}

	for i := 0; i < length; i++{
		minIndex := i
		for j := i+1; j < length; j++{
			if sli[minIndex] < sli[j] {
				minIndex = j
			}
		}

		sli[i], sli[minIndex] = sli[minIndex], sli[i]
	}
	return sli
}