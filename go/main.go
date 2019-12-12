package main

import "fmt"

//Examples:
//%f     default width, default precision
//%9f    width 9, default precision
//%.2f   default width, precision 2
//%9.2f  width 9, precision 2
//%9.f   width 9, precision 0

//如下示例
func main() {
	fmt.Printf("%b\n", 10.45)    //5882827013252710p-49
	fmt.Printf("%e\n", 10.45)    //1.045000E+01
	fmt.Printf("%E\n", 10.45)    //1.045000E+01
	fmt.Printf("%f\n", 10.451)   //10.451000
	fmt.Printf("%9f\n", 10.45)   //10.450000
	fmt.Printf("%9.2f\n", 10.45) //     10.45
	fmt.Printf("%9.f\n", 10.45)  //        10
	fmt.Printf("%.1f\n", 10.45)  //10.4
	fmt.Printf("%.4f\n", 10.45)  //10.4500
	fmt.Printf("%F\n", 10.45)    //10.450000
	fmt.Printf("%g\n", 10.45)    //10.45
	fmt.Printf("%G\n", 10.45)    //10.45
	fmt.Printf("%x\n", 10.45)    //0x1.4e66666666666p+03
	fmt.Printf("%X\n", 10.45)    //0X1.4E66666666666P+03
}
