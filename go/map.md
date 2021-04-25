# 关于 Map
<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/go/map.jpg'>
</p>

# map 设计中常用的数据结构
- 哈希查找表: hash函数+数组, 最差时间复杂度O(N),通常是O(1)
- 搜索树: 自平衡搜索树，通常使用AVL树、红黑树，最差时间复杂度O(logn)

# 哈希冲突
不同的key经过hash函数运算后分配到了同一个bucket，产生冲突
## 哈希冲突解决方式
- 拉链法: 将bucket设计成一个链表，每一个bucket会记录下一个节点地址
- 开放寻址法: 在当前bucket数组中往后寻找空位，如果后面没有从开头开始寻找

# 底层结构 (runtime包map.go文件中)
```
// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	// 调用len()函数返回的元素个数
	count     int // # live cells == size of map.  Must be first (used by len() builtin) 
	flags     uint8    // flags
	// buckets 数组的长度的对数
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)  
	// 溢出bucket个数
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details 
	// hash种子  
	hash0     uint32 // hash seed              
    // 指向 buckets 数组，大小为 2^B, 如果元素个数为0，就为 nil
	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.  
	// 结构扩容的时候用于复制的buckets数组指针
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing  
	// 指示扩容进度，小于此地址的 buckets 迁移完成
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated) 

	extra *mapextra // optional fields
}

// mapextra holds fields that are not present on all maps.
type mapextra struct {
	// If both key and value do not contain pointers and are inline, then we mark bucket
	// type as containing no pointers. This avoids scanning such maps.
	// However, bmap.overflow is a pointer. In order to keep overflow buckets
	// alive, we store pointers to all overflow buckets in hmap.extra.overflow and hmap.extra.oldoverflow.
	// overflow and oldoverflow are only used if key and value do not contain pointers.
	// overflow contains overflow buckets for hmap.buckets.
	// oldoverflow contains overflow buckets for hmap.oldbuckets.
	// The indirection allows to store a pointer to the slice in hiter.
	// hmap.buckets （当前）溢出桶的指针地址
	overflow    *[]*bmap  
	// hmap.oldbuckets （旧）溢出桶的指针地址
	oldoverflow *[]*bmap

	// nextOverflow holds a pointer to a free overflow bucket.
	// 为空闲溢出桶的指针地址
	nextOverflow *bmap
}

// A bucket for a Go map.
type bmap struct {
	// tophash generally contains the top byte of the hash value
	// for each key in this bucket. If tophash[0] < minTopHash,
	// tophash[0] is a bucket evacuation state instead.
	tophash [bucketCnt]uint8     // 存储哈希值的高8位
	// Followed by bucketCnt keys and then bucketCnt values.
	// NOTE: packing all the keys together and then all the values together makes the
	// code a bit more complicated than alternating key/value/key/value/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
}
```

# hash 函数
go中map关于hash函数的选择，程序启动时会检测cpu是否支持aes, 如果支持则使用aes hash, 如果不支持则使用memhash, 这是在函数 alginit() 中完成，位于路径：src/runtime/alg.go


# map结构图

![map结构示意图](https://github.com/w1991668899/blog/blob/master/image/go/map%E7%BB%93%E6%9E%84%E7%A4%BA%E6%84%8F%E5%9B%BE.png)

**存储形式:** kv的存储形式为 k0k1k2k3…k7v0v1v2…v7,这样做的好处是：在key和value的长度不同的时候，节省padding空间。
# 负载因子

> 负载因子 = 键数量/bucket数量

负载因子需要控制在合适的大小，超过其阀值需要进行rehash，即键值对重新组织.

- 哈希因子过小空间利用率低
- 哈希因子过大，键冲突严重，存取效率低

**查找过程:** 

- 根据key算出16位hash值
- hash值低8位与`hmpa.B`取模确定bucket位置
- hash值高8位在`tophash`数组中
- 如果`tophash[i]`中存储值也哈希值相等，则去找到该bucket中的key值进行比较
- 当前bucket没有找到，则继续从下个overflow的bucket中查找
- 如果当前处于搬迁过程，则优先从oldbuckets查找

**插入过程**

- 根据key算出16位hash值
- hash值低8位与`hmpa.B`取模确定bucket位置
- hash值高8位在`tophash`数组中
- 查找该key是否已经存在，如果存在则直接更新值
- 如果没找到将key，将key插入

**删除**

- 如果`key`是一个指针类型的，则直接将其置为空，等待GC清除；
- 如果是值类型的，则清除相关内存。
- 同理，对``value``做相同的操作。
- 最后把key对应的高位值对应的数组index置为空。

# map中的key为什么是无序的？
- map遍历的时候并不是从第0个bucket开始的，每次都是随机一个开始
- map扩容的时候数据是分部迁移的


# 删除过程？

# 遍历过程？ 

# 赋值过程？

# 扩容过程？

# 

# map中的key
```
package main

import "fmt"

func main()  {
	var badMap2 = map[interface{}]int{
		0:1000,
		"1":1000,
		[2]int{10}:1000,
	}
	fmt.Println(badMap2)
}

// 输出:map[[10 0]:1000 1:1000 0:1000]
```

- map中的key不能是 slice、function、map, 这些类型只能跟nil做`==`判断<br>

**chan类型可以作为key但是后面的值会覆盖前面的**
```
package main

import "fmt"

func main()  {
	ch:= make(chan string, 2)
	ma := make(map[chan string]int)

	ch<- "study"
	ma[ch] = 100
	ch<- "go"
	ma[ch] = 999

	fmt.Println(ma)
}

// 输出: map[0xc00008a0c0:999]

```

**uintptr也可以作为key**
```
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

// 输出:
// map[100:study 200:go]
// map[100:study 200:go 824634257240:together]
```





