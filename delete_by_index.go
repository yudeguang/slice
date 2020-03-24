// DeleteByIndexright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a DeleteByIndex of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

/*
// 第一种方法（保持剩余元素的次序）：
s = append(s[:from], s[to:]...)

// 第二种方法（保持剩余元素的次序）：
s = s[:from + copy(s[from:], s[to:])]

// 第三种方法（不保持剩余元素的次序）：
if n := to-from; len(s)-to < n {
	copy(s[from:to], s[to:])
} else {
	copy(s[from:to], s[len(s)-n:])
}
s = s[:len(s)-(to-from)]
//如果切片的元素可能引用着其它值，则我们应该重置因为删除元素而多出来的元素槽位上的元素值，以避免暂时性的内存泄露：
// "len(s)+to-from"是删除操作之前切片s的长度。
temp := s[len(s):len(s)+to-from]
for i := range temp {
	temp[i] = t0 // t0是类型T的零值
}

//在此，只实现第一种方法，并且使用安全一些的方法，保持原有切片内容不变

*/

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndex(s []interface{}, from int, to ...int) []interface{} {
	copyS := Copy(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexString(s []string, from int, to ...int) []string {
	copyS := CopyString(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexInt(s []int, from int, to ...int) []int {
	copyS := CopyInt(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexInt8(s []int8, from int, to ...int) []int8 {
	copyS := CopyInt8(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexInt16(s []int16, from int, to ...int) []int16 {
	copyS := CopyInt16(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexInt32(s []int32, from int, to ...int) []int32 {
	copyS := CopyInt32(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexInt64(s []int64, from int, to ...int) []int64 {
	copyS := CopyInt64(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexUint(s []uint, from int, to ...int) []uint {
	copyS := CopyUint(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexUint8(s []uint8, from int, to ...int) []uint8 {
	copyS := CopyUint8(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexUint16(s []uint16, from int, to ...int) []uint16 {
	copyS := CopyUint16(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexUint32(s []uint32, from int, to ...int) []uint32 {
	copyS := CopyUint32(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexUint64(s []uint64, from int, to ...int) []uint64 {
	copyS := CopyUint64(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexUintptr(s []uintptr, from int, to ...int) []uintptr {
	copyS := CopyUintptr(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexFloat32(s []float32, from int, to ...int) []float32 {
	copyS := CopyFloat32(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexFloat64(s []float64, from int, to ...int) []float64 {
	copyS := CopyFloat64(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexRune(s []rune, from int, to ...int) []rune {
	copyS := CopyRune(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexByte(s []byte, from int, to ...int) []byte {
	copyS := CopyByte(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexComplex64(s []complex64, from int, to ...int) []complex64 {
	copyS := CopyComplex64(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexComplex128(s []complex128, from int, to ...int) []complex128 {
	copyS := CopyComplex128(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}

//删除从位置是从from（包括）到to（不包括）的相关元素,如果to不填写，则表示只删除from这个位置的一个元素
func DeleteByIndexBool(s []bool, from int, to ...int) []bool {
	copyS := CopyBool(s)
	if len(to) == 0 {
		return append(copyS[:from], copyS[from+1:]...)
	}
	return append(copyS[:from], copyS[to[0]:]...)
}
