# channel

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/go/channel.jpeg'>
</p>

channel提供了一种机制。它既可以同步两个并发执行的函数，又可以让两个函数通过相互传递特定类型的值来通信。

# 类型声明

```
chan T   

type StrChan chan string   // 自定义类型

type IntChan = chan int    // 别名

var ch chan int            //channel的零值为：nil
```

# 初始化

```
ch1 := make(chan int)       // 无缓冲 channel
ch2 := make(chan int, 0)    // 无缓冲 channel
ch3 := make(chan int, 5)    // 缓冲 channel
```

注意:<br><br>
一个通道值的缓冲容量总是不变的。如果第二个参数省略，就表示该通道永远无法缓冲任何元素，发送给它的元素应该立即被取走，否则发送方的的 `goroutine` 将被阻塞，直到元素被接收.

# 关闭通道

```
strChan := make(chan string, 5)  

close(strChan)   // 关闭通道
```
注意:<br><br>

- 向一个已经关闭的通道发送数据会引起运行时恐慌，所以请确保在发送端关闭通道
- 当通道被关闭后接收端仍然可以接收(包括缓冲通道与无缓冲通道), 如果通道已经没有元素则接收到的值为通道元素零值，可通过 `ok` 返回值值判断
- 对通道的重复 `close` 将引起运行时恐慌
- 执行 `close` 关闭通道必须传入通道变量,否则默认为传入 `nil` 将引起运行时恐慌

# 发送与接收元素

```
strChan := make(chan string, 5)     //初始化

strChan<- "shanghai"            //发送元素
strChan<- "beijing"             

elem := <-strChan               // 接收元素
elem, ok := <-strChan           // 接收并判断是否接收成功
```

注意:<br><br>
- 因为元素的零值也可以发送到通道中，所以当接收到这样一个值的时候我们可能无法判断该元素是正确的还是该通道已经关闭，我们接收到的是关闭后的零值，这时候通过 `ok` 判定
当从一个已经 ``

# 关于值的传递

- 经由通道传递的值至少会被复制一次，至多复制两次
- 当向一个已空通道发送值且已有至少一个接收方在等待时，该通道会跳过本身的缓冲队列，直接把值复制给接收方
- 通道属于环形队列当从一个已满的通道接收元素且已经有至少一个发送方在等待时，通道会把缓冲队列中最早进入的那个元素复制给对方，再把最早等待的发送方的数据复制到原先位置上
- 当通道中的元素是值类型时不会影响到发送方，当包含引用类型是对元素的修改会影响发送与接收两端


# 单向 `channel`

```
var sendCh chan<- int       // 发送通道
var receiveCh  <-chan int   // 接收通道
```
注意:<br><br>
函数声明中双向通道可以转为单向通道，这是go的一个语法糖，但不可反过来

# `for` `select`

```
package main

import "fmt"

func main() {
	intChan := make(chan int, 10)

	for i := 0; i < 10; i++{
		intChan<- i
	}

	close(intChan)

	syncChan := make(chan struct{}, 1)

	go func() {
		Loop:
			for{
				select {
				case e, ok := <-intChan:
					if !ok {
						fmt.Println("End.")
						break Loop
					}
					fmt.Printf("Received: %v\n", e)
				}
			}
		syncChan<- struct{}{}
	}()
	<-syncChan
}

// 打印结果
Received: 0
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
Received: 6
Received: 7
Received: 8
Received: 9
End.

```







