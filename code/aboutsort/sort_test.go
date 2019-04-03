package aboutsort

import (
	"fmt"
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	var sli = []int{1,50,999,6,6,3,7,5,10}
	fmt.Println("冒泡排序: 小=>大")
	fmt.Println("排序前:", sli)
	BubbleSort1(sli)
	fmt.Println("排序后:", sli)
}

func TestBubbleSort2(t *testing.T) {
	var sli = []int{1,50,999,6,6,3,7,5,10}
	fmt.Println("冒泡排序: 大=>小")
	fmt.Println("排序前:", sli)
	BubbleSort2(sli)
	fmt.Println("排序后:", sli)
}

func TestInsertSort1(t *testing.T) {
	var sli = []int{1,50,999,6,6,3,7,5,10}
	fmt.Println("插入排序: 小=>大")
	fmt.Println("排序前：", sli)
	InsertSort1(sli)
	fmt.Println("排序后：", sli)
}

func TestInsertSort2(t *testing.T) {
	var sli = []int{1,50,999,6,6,3,7,5,10}
	fmt.Println("插入排序: 大=>小")
	fmt.Println("排序前：", sli)
	InsertSort2(sli)
	fmt.Println("排序后：", sli)
}

func TestSelectSort(t *testing.T) {
	var sli = []int{1,50,999,6,6,3,7,5,10}
	fmt.Println("选择排序: 小=>大")
	fmt.Println("排序前：", sli)
	SelectSort1(sli)
	fmt.Println("排序后：", sli)
}

func TestSelectSort2(t *testing.T) {
	var sli = []int{1,50,999,6,6,3,7,5,10}
	fmt.Println("选择排序: 大=>小")
	fmt.Println("排序前：", sli)
	SelectSort2(sli)
	fmt.Println("排序后：", sli)
}

func TestMergeSort(t *testing.T) {
	var sli = []int{1,50,999,6,6,3,7,5,10}
	fmt.Println("归并排序: 小=>大")
	fmt.Println("排序前：", sli)
	MergeSort(sli)
	fmt.Println("排序后：", sli)
}



