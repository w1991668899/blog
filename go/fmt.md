# Fmt 格式化

## [参考](https://pkg.go.dev/fmt?tab=doc)

通用verbs
====================
```
%v	     　值的默认格式
%+v      　类似%v,但输出结构体时会添加字段名
%#v	　　 Go语法表示值
%T	　　 Go语法表示类型
%%	   　 百分号表示

//如下示例：
type Sample struct {
	Title   string
	name 	string
	Age     int
}

func main() {

	s := Sample{"测试", "wentao", 26}

	fmt.Printf("%v\n", s)    // {测试 wentao 26}
	fmt.Printf("%+v\n", s) 	// {Title:测试 name:wentao Age:26}
	fmt.Printf("%#v\n", s) 	// main.Sample{Title:"测试", name:"wentao", Age:26}
	fmt.Printf("%T\n", s)   	//  main.Sample
	fmt.Printf("%v%%\n", s.Age) //  26%
}

```
布尔值
============================================
```
%t	true或false 
//如下示例
func main() {
	fmt.Printf("%t\n", true)  //true  注: 必须为Bool类型才能使用 %t 占位符
	fmt.Printf("%t\n", "string")        //%!t(string=string)   错误
	fmt.Printf("%t\n", 100)     //%!t(int=100)  错误
	fmt.Printf("%t\n", 0)   //%!t(int=0)  错误
}
```
整数
============================================
```
%b  表示二进制
%c  相应的Unicode代码点表示的字符
%d  表示十进制
%o  表示八进制
%O	0o 前缀表示的八进制
%q  该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x  表示为十六进制，使用a-f
%X  表示为十六进制，使用A-F
%U  表示为Unicode格式：U+1234，等价于"U+%04X"

//如下示例
func main() {
	fmt.Printf("%b\n", 26)			//11010
	fmt.Printf("%c\n", 0x4E2D)		//中
	fmt.Printf("%d\n", 0x12)		//18
	fmt.Printf("%o\n", 20)			//24
	fmt.Printf("%O\n", 20)			//0o24
	fmt.Printf("%q\n", 0x4E2D)		//'中'
	fmt.Printf("%x\n", 14)			//e
	fmt.Printf("%X\n", 14)			//E
	fmt.Printf("%U\n", 0x4E2D)     //U+4E2D
}
```

拓展
================================================
```
package main

import "fmt"

//+    总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
//-    在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
//#    切换格式：
//八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）；
//对%q（%#q），如果strconv.CanBackquote返回真会输出反引号括起来的未转义字符串；
//对%U（%#U），如果字符是可打印的，会在输出Unicode格式、空格、单引号括起来的go字面值；

//' '    对数值，正数前加空格而负数前加负号；
//对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格；
//0    使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；

func main() {
	fmt.Printf("%o\n", 100)     //144
	fmt.Printf("%#o\n", 100)    //0144

	fmt.Printf("%x\n", 100)     //64
	fmt.Printf("%#x\n", 100)    //0x64

	fmt.Printf("%p\n", new(int))    //0xc00009a010
	fmt.Printf("%#p\n", new(int))   //c00009a018

	fmt.Printf("%U\n", 'G')     //U+0047
	fmt.Printf("%#U\n", 'G')    //U+0047 'G'

	fmt.Printf("% d", 100)    // 100
	fmt.Printf("% x", 100)      // 64
}
```

按索引取值
===============================================
```
package main

import "fmt"

//紧跟在verb之前的[n]符号表示应格式化第n个参数（索引从1开始）。同样的在'*'之前的[n]符号表示采用第n个参数的值作为宽度或精度。在处理完方括号表达式[n]后，除非另有指示，会接着处理参数n+1，n+2……（就是说移动了当前处理位置）
func main() {
	fmt.Printf("%[2]d %[1]d\n", 11, 22)             //22 11
	fmt.Printf("%[3]*.[2]*[1]f\n", 12.0, 2, 6)      // 12.00
	fmt.Printf("%d %d %#[1]x %#x\n", 16, 17)        //16 17 0x10 0x11
}
```

浮点数与复数
================================================
```
%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
//无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat %e    科学计数法，如-1234.456e+78 %E    科学计数法，如-1234.456E+78 %f    有小数部分但无指数部分，如123.456 %F    等价于%f %g    根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）

%e	科学计数法，例如 -1234.456e+78 
%E	科学计数法，例如 -1234.456E+78
%f	有小数点而无指数，例如 123.456 
%F	等价于%f
%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
%x	小写16进制
%X	大写16进制

Examples:
		%f     default width, default precision
		%9f    width 9, default precision
		%.2f   default width, precision 2  默认宽度，精度2， 默认向右对齐, 左边补空格
		%9.2f  width 9, precision 2,  宽度9，精度2， 默认向右对齐, 左边补空格
		%9.f   width 9, precision 0, 宽度9，默认精度， 默认向右对齐, 左边补空格

//如下示例
func main() {
	fmt.Printf("%b\n", 10.45)		//5882827013252710p-49
	fmt.Printf("%e\n", 10.45)		//1.045000E+01
	fmt.Printf("%E\n", 10.45)		//1.045000E+01
	fmt.Printf("%f\n", 10.451)		//10.451000
	fmt.Printf("%9f\n", 10.45)		//10.450000
	fmt.Printf("%9.2f\n", 10.45)	//     10.45, 宽度9，精度2， 默认向右对齐
	fmt.Printf("%9.f\n", 10.45)		//        10
	fmt.Printf("%.1f\n", 10.45)		//10.4
	fmt.Printf("%.4f\n", 10.45)		//10.4500
	fmt.Printf("%F\n", 10.45)		//10.450000
	fmt.Printf("%g\n", 10.45)		//10.45
	fmt.Printf("%G\n", 10.45)		//10.45
	fmt.Printf("%x\n", 10.45)		//0x1.4e66666666666p+03
	fmt.Printf("%X\n", 10.45)		//0X1.4E66666666666P+03
}

```
string与[]byte
================================================================
```
%s	输出字符串表示（string类型或[]byte) 
%q	双引号围绕的字符串，由Go语法安全地转义
%x	十六进制，小写字母，每字节两个字符 （使用a-f）
%X	十六进制，大写字母，每字节两个字符 （使用A-F） 

//如下示例：
package main

import "fmt"

func main() {
	fmt.Printf("%s\n", []byte("go开发"))		//go开发
	fmt.Printf("%s\n", "go开发")			//go开发
	fmt.Printf("%q\n", "go开发")			//"go开发"
	fmt.Printf("%x\n", "go开发")			//676fe5bc80e58f91
	fmt.Printf("%X\n", "go开发")			//676FE5BC80E58F91
}
```
Slice
===================================================================
```
%p       切片第一个元素的指针

The %b, %d, %o, %x and %X verbs also work with pointers,
formatting the value exactly as if it were an integer.

//如下示例
package main

import "fmt"

func main() {
	fmt.Printf("%p\n", []byte("go开发"))			//0xc42001a0d8
	fmt.Printf("%p\n", []int{1, 2, 3, 45, 65})	//0xc420020180
}
```
point
====================================================================
```
%p       十六进制内存地址,前缀ox
package main

import "fmt"

func main() {

	str := "go开发"
	fmt.Printf("%p\n", &str)			//0xc42000e1e0
}
```
说明：
1. Go没有 '%u' 点位符,整数如果是无符号类型自然输出也是无符号的。类似的，也没有必要指定操作数的尺寸（int8，int64）。
2. 宽度与精度的控制格式以Unicode码点为单位(不同于C的printf，它的这两个因数指的是字节的数量)。宽度为该数值占用区域的最小宽度；精度为小数点之后的位数。操作数的类型为int时，宽度与精度都可用字符 '*' 表示。
3. 对于 %g/%G 而言，精度为所有数字的总数，例如：123.45，%.4g 会打印123.5，（而 %6.2f 会打印123.45）。%e 和 %f 的默认精度为6
4. 对大多数的数值类型而言，宽度为输出的最小字符数，如果必要的话会为已格式化的形式填充空格。而以字符串类型，精度为输出的最大字符数，如果必要的话会直接截断。
5. 对复数，宽度和精度会分别用于实部和虚部，结果用小括号包裹。因此%f用于1.2+3.4i输出(1.200000+3.400000i)。
6.
 '+'	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
' '	对数值，正数前加空格而负数前加负号；
'-'	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
'#'	切换格式：
  	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）；
 	对%q（%#q），如果strconv.CanBackquote返回真会输出反引号括起来的未转义字符串；
 	对%U（%#U），输出Unicode格式后，如字符可打印，还会输出空格和单引号括起来的go字面值；
  	对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格；
'0'	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；

待续。。。







