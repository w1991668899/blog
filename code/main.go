package main

import (
	"code/aboutsort"
	"fmt"
)

func main()  {
	var sli = []int{1,50,999,6,6,3,7,5,10}
	fmt.Println("插入排序: 小=>大")
	fmt.Println("排序前：", sli)
	aboutsort.InsertSort1(sli)
	fmt.Println("排序后：", sli)
}
