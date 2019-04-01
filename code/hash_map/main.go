package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	ma := make(map[uintptr]string)

	a := 100
	pt := uintptr(unsafe.Pointer(&a))

	ma[100] = "study"
	ma[200] = "go"
	fmt.Println(ma)

	ma[pt] = "together"

	fmt.Println(ma)
}
