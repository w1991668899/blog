# 关于 Map
<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/go/map.jpg'>
</p>

# 底层结构 (runtime包map.go文件中)
```
// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin) 元素个数
	flags     uint8    // flags
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)  2^B个bucket
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details  溢出bucket个数
	hash0     uint32 // hash seed   hash种子             

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.  buckets的数组指针
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing  结构扩容的时候用于复制的buckets数组指针
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated) 搬迁进度（已经搬迁的buckets数量）

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
	overflow    *[]*bmap   // 溢出bucket的地址
	oldoverflow *[]*bmap

	// nextOverflow holds a pointer to a free overflow bucket.
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




