包([github.com/yudeguang/slice](http://github.com/yudeguang/slice "github.com/yudeguang/slice"))，用于切片处理时经常会遇到的一些操作,主要包括如下函数：

- **Copy** 安全复制

- **Contains** 判断是否包含

- **InsertIgnore** 插入不重复元素

- **DeleteByElement** 根据元素删除(不会改变原有切片)
- **DeleteByIndex** 根据下标范围删除，也可以只删除单个下标位置的元素(不会改变原有切片)

- **Distinct** 去重复(元素个数较大时会自动切换为hash算法)

- **InnerJoin** 内连接(元素个数较大时会自动切换为hash算法)
- **Union** 合并去重(元素个数较大时会自动切换为hash算法)
- **NotIn** 相减--返回mNotInN(元素个数较大时会自动切换为hash算法)

- Left 返回左侧若干元素
- Right 返回右侧若干元素

- **Reverse** 倒序排列
- **Shuffling** 随机
- **Sort** 排序

由于go暂时不支持泛型，所以目前是针对每一种基础类型(**int,string等**)都分别编写了所对应的相关函数。以string类型为列:

```go
package main

import (
	"fmt"
	"github.com/yudeguang/slice"
	"github.com/yudeguang/stringsx"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	s1 := "abcdefabcdef3434"
	s2 := "bcdefgbcdefg02323"
	m := stringsx.SplitByLen(s1, 1)
	n := stringsx.SplitByLen(s2, 1)
	//先观察一下两个切片
	log.Println("原始m,n分别如下:")
	log.Println("m为:")
	log.Println(m)
	log.Println("n为:")
	log.Println(n)

	//复制
	fmt.Println("")
	log.Println("复制CopyString:")
	copyM := slice.CopyString(m)
	log.Println("copyM为:")
	log.Println(copyM)
	log.Println("m为:")
	log.Println(m)
	log.Printf("copyM的地址:%p", copyM)
	log.Printf("    m的地址:%p", m)
	/*
			增删改查
			增:实现了InsertIgnore
		    删:实现以下两种,根据元素值删除DeleteByElement,根据下标索引位置删除DeleteByIndex
	*/
	//插入不重复数据
	fmt.Println("")
	log.Println("插入不重复数据InsertIgnoreString:")
	log.Println("m插入元素y后:")
	log.Println(slice.InsertIgnoreString(m, "y"))
	log.Println("m为:")
	log.Println(m)
	log.Println("m插入元素a后:")
	log.Println(slice.InsertIgnoreString(m, "a"))
	log.Println("m为:")
	log.Println(m)

	//删除某个元素
	fmt.Println("")
	log.Println("删除元素DeleteByElementString:")
	log.Println("m删除元素a后的结果为:")
	log.Println(slice.DeleteByElementString(m, "a"))
	log.Println("m为:")
	log.Println(m)

	//删除第i个元素
	fmt.Println("")
	log.Println("删除第i个元素DeleteByIndexString:")
	log.Println("m删除下标为3的元素后的结果为:")
	log.Println(slice.DeleteByIndexString(m, 3))
	log.Println("m为:")
	log.Println(m)

	//删除一段元素
	fmt.Println("")
	log.Println("删除一段元素DeleteByIndexString:")
	log.Println("m删除下标从3到5(不包含5)的元素后的结果为:")
	log.Println(slice.DeleteByIndexString(m, 3, 5))
	log.Println("m为:")
	log.Println(m)

	//是否包含
	fmt.Println("")
	log.Println("是否包含ContainsString:")
	log.Println("m是否包含元素a", slice.ContainsString(m, "a"))
	log.Println("m是否包含元素x", slice.ContainsString(m, "x"))
	log.Println("m为:")
	log.Println(m)

	//去重复
	fmt.Println("")
	log.Println("去重复DistinctString:")
	log.Println("m去重复后的结果为:")
	log.Println(slice.DistinctString(m))
	log.Println("m为:")
	log.Println(m)
	/*
		集合操作，实现了:
		innerJoin
		union
		notin
	*/
	//内连接
	fmt.Println("")
	log.Println("内连接innerJoinString:")
	log.Println("m与n内连接结果为:")
	log.Println(slice.InnerJoinString(m, n))
	log.Println("m为:")
	log.Println(m)
	log.Println("n为:")
	log.Println(n)

	//合并去重
	fmt.Println("")
	log.Println("合并去重UnionString:")
	log.Println("m与n合并去重结果为:")
	log.Println(slice.UnionString(m, n))
	log.Println("m为:")
	log.Println(m)
	log.Println("n为:")
	log.Println(n)

	//集合m不在集合n中的元素即m-n 注意，不对m做去重操作
	fmt.Println("")
	log.Println("相减NotInString:")
	log.Println("m-n结果为:")
	log.Println(slice.NotInString(m, n))
	log.Println("m为:")
	log.Println(m)
	log.Println("n为:")
	log.Println(n)
	/*
		以下操作会改变切片本身
		Reverse
		sort
		Shuffling
	*/
	//倒序
	fmt.Println("")
	log.Println("倒序ReverseString:")
	slice.ReverseString(m)
	log.Println("m为:")
	log.Println(m)

	//从小到大排序
	fmt.Println("")
	log.Println("排序SortString:")
	slice.SortString(m)
	log.Println("m为:")
	log.Println(m)

	//随机打乱
	fmt.Println("")
	log.Println("随机打乱ShufflingString:")
	slice.ShufflingString(m)
	log.Println("m为:")
	log.Println(m)
	slice.ShufflingString(m)
	log.Println("m为:")
	log.Println(m)
}
```

执行结果如下：

```go
22:32:41 main.go:17: 原始m,n分别如下:
22:32:41 main.go:18: m为:
22:32:41 main.go:19: [a b c d e f a b c d e f 3 4 3 4]
22:32:41 main.go:20: n为:
22:32:41 main.go:21: [b c d e f g b c d e f g 0 2 3 2 3]

22:32:41 main.go:25: 复制CopyString:
22:32:41 main.go:27: copyM为:
22:32:41 main.go:28: [a b c d e f a b c d e f 3 4 3 4]
22:32:41 main.go:29: m为:
22:32:41 main.go:30: [a b c d e f a b c d e f 3 4 3 4]
22:32:41 main.go:31: copyM的地址:0xc042086100
22:32:41 main.go:32:     m的地址:0xc042086000

22:32:41 main.go:40: 插入不重复数据InsertIgnoreString:
22:32:41 main.go:41: m插入元素y后:
22:32:41 main.go:42: [a b c d e f a b c d e f 3 4 3 4 y]
22:32:41 main.go:43: m为:
22:32:41 main.go:44: [a b c d e f a b c d e f 3 4 3 4]
22:32:41 main.go:45: m插入元素a后:
22:32:41 main.go:46: [a b c d e f a b c d e f 3 4 3 4]
22:32:41 main.go:47: m为:
22:32:41 main.go:48: [a b c d e f a b c d e f 3 4 3 4]

22:32:41 main.go:52: 删除元素DeleteByElementString:
22:32:41 main.go:53: m删除元素a后的结果为:
22:32:41 main.go:54: [b c d e f b c d e f 3 4 3 4]
22:32:41 main.go:55: m为:
22:32:41 main.go:56: [a b c d e f a b c d e f 3 4 3 4]

22:32:41 main.go:60: 删除第i个元素DeleteByIndexString:
22:32:41 main.go:61: m删除下标为3的元素后的结果为:
22:32:41 main.go:62: [a b c e f a b c d e f 3 4 3 4]
22:32:41 main.go:63: m为:
22:32:41 main.go:64: [a b c d e f a b c d e f 3 4 3 4]

22:32:41 main.go:68: 删除一段元素DeleteByIndexString:
22:32:41 main.go:69: m删除下标从3到5(不包含5)的元素后的结果为:
22:32:41 main.go:70: [a b c f a b c d e f 3 4 3 4]
22:32:41 main.go:71: m为:
22:32:41 main.go:72: [a b c d e f a b c d e f 3 4 3 4]

22:32:41 main.go:76: 是否包含ContainsString:
22:32:41 main.go:77: m是否包含元素a true
22:32:41 main.go:78: m是否包含元素x false
22:32:41 main.go:79: m为:
22:32:41 main.go:80: [a b c d e f a b c d e f 3 4 3 4]

22:32:41 main.go:84: 去重复DistinctString:
22:32:41 main.go:85: m去重复后的结果为:
22:32:41 main.go:86: [a b c d e f 3 4]
22:32:41 main.go:87: m为:
22:32:41 main.go:88: [a b c d e f a b c d e f 3 4 3 4]

22:32:41 main.go:97: 内连接innerJoinString:
22:32:41 main.go:98: m与n内连接结果为:
22:32:41 main.go:99: [b c d e f 3]
22:32:41 main.go:100: m为:
22:32:41 main.go:101: [a b c d e f a b c d e f 3 4 3 4]
22:32:41 main.go:102: n为:
22:32:41 main.go:103: [b c d e f g b c d e f g 0 2 3 2 3]

22:32:41 main.go:107: 合并去重UnionString:
22:32:41 main.go:108: m与n合并去重结果为:
22:32:41 main.go:109: [a b c d e f 3 4 g 0 2]
22:32:41 main.go:110: m为:
22:32:41 main.go:111: [a b c d e f a b c d e f 3 4 3 4]
22:32:41 main.go:112: n为:
22:32:41 main.go:113: [b c d e f g b c d e f g 0 2 3 2 3]

22:32:41 main.go:117: 相减NotInString:
22:32:41 main.go:118: m-n结果为:
22:32:41 main.go:119: [a a 4 4]
22:32:41 main.go:120: m为:
22:32:41 main.go:121: [a b c d e f a b c d e f 3 4 3 4]
22:32:41 main.go:122: n为:
22:32:41 main.go:123: [b c d e f g b c d e f g 0 2 3 2 3]

22:32:41 main.go:132: 倒序ReverseString:
22:32:41 main.go:134: m为:
22:32:41 main.go:135: [4 3 4 3 f e d c b a f e d c b a]

22:32:41 main.go:139: 排序SortString:
22:32:41 main.go:141: m为:
22:32:41 main.go:142: [3 3 4 4 a a b b c c d d e e f f]

22:32:41 main.go:146: 随机打乱ShufflingString:
22:32:41 main.go:148: m为:
22:32:41 main.go:149: [b c 4 3 4 e c a 3 a e b f d f d]
22:32:41 main.go:151: m为:
22:32:41 main.go:152: [a a 4 b 3 f 3 e c 4 d c d b f e]
```