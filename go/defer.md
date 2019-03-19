# 关于 defer

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/go/defer.jpeg'>
</p>

#什么是defer
A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a [return statement](https://golang.org/ref/spec#Return_statements), reached the end of its [function body](https://golang.org/ref/spec#Function_declarations), or because the corresponding goroutine is [panicking](https://golang.org/ref/spec#Handling_panics).

通过使用 `defer` 修饰一个函数，使其在外部函数返回后才被执行，一种是因为周围的函数执行了一个 `return`语句，到达了它的函数体的末尾，另一种是因为相应的goroutine是 panicking

#先看example1
```
package main

import "fmt"

func main()  {
	data1 := example1()
	fmt.Printf("output1:%v\n", data1)
	data2 := example2()
	fmt.Printf("output2:%v\n", data2)
	data3 := example3()
	fmt.Printf("output3:%v\n", data3)
}

func example1() (result int) {
	defer func() {
		result++
	}()
	return 0
}
func example2() (r int) {
	t := 1
	defer func() {
		t = t + 1
	}()
	return t
}
func example3() (r int) {
	defer func(r int) {
		r = r + 1
	}(r)
	return 1
}
//output1:1
//output2:1
//output3:1
```

#go 函数返回值的栈空间与执行顺序
1. golang 的返回值是通过栈空间，不是通过寄存器，这点最重要。调用函数前，首先分配的是返回值空间，然后是入参地址，再是其他临时变量地址。
2. return 是非原子的，return 操作的确是分三步:
&emsp;将返回值拷贝到栈空间第一块区域，即：返回值 = x
&emsp;执行defer 操作, 即: 调用defer函数
&emsp;RET 跳转， 即：返回
这个操作的确是可以对应到汇编代码的。

#example4
```
package main

import "fmt"

type message struct {
	content string
}

func (p *message) set(c string) {
	p.content = c
}

func (p *message) print() string {
	return p.content
}

func main()  {

	m := &message{content: "Hello"}

	defer fmt.Println(m.print())        //m.print()  == "Hello", 将值拷贝进去

	m.set("World")
}

//输出：Hello
```
为什么输出的是"Hello": 在 defer 中， fmt.Print 被推迟到函数返回后执行，可是 m.print() 的值在当时就已经被求出，因此， m.print() 会返回 "Hello" ，这个值会保存一直到外围函数返回。

#example4 改进
```
package main

import "fmt"

type message struct {
	content string
}

func (p *message) set(c string) {
	p.content = c
}

func (p *message) print() string {
	return p.content
}

func main()  {

	m := &message{content: "Hello"}

	defer func() {
		fmt.Println(m.print())
	}()

	m.set("World")
}

//输出:World
```
对 m 地址的引用

#example 5
```
package main

import (
	"errors"
	"fmt"
	"io"
)

type reader struct{}

func (r reader) Close() error {
	return errors.New("Close Error")
}


func release(r io.Closer) (err error)  {

	defer func() {
		if err := r.Close(); err != nil{
			fmt.Println("my errors")
		}
	}()

	return
}

func main()  {

	r := reader{}

	err := release(r)

	fmt.Println(err)
}
//输出：my errors
//输出：<nil>
```
#example 5改进
```
package main

import (
	"errors"
	"fmt"
	"io"
)

type reader struct{}

func (r reader) Close() error {
	return errors.New("Close Error")
}


func release(r io.Closer) (err error)  {

	defer func() {
		if err = r.Close(); err != nil{
			fmt.Println("my errors")
		}
	}()

	return
}

func main()  {

	r := reader{}

	err := release(r)

	fmt.Println(err)
}

//输出: my errors
//输出：Close Error
```
因为变量的作用域不同导致的块级屏问题。将 err := r.Close() 改成 err = r.Close 即可。

#example 6   忽略了错误
```
func do() error {
    f, err := os.Open("book.txt")
    if err != nil {
        return err
    }
    defer f.Close()

    // ..code...

    return nil
}
```
#example 6 改进
```
func do() (err error) {
    f, err := os.Open("book.txt")
    if err != nil {
        return err
    }

    defer func() {
        if ferr := f.Close(); ferr != nil {
            err = ferr
        }
    }()

    // ..code...

    return nil
}
```








