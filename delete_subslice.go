// DeleteSubSliceright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a DeleteSubSlice of the MIT was not distributed with this file,
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

//在此，只实现第一种方法

*/
//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSlice(s []interface{}, from, to int) []interface{} {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceString(s []string, from, to int) []string {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceInt(s []int, from, to int) []int {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceInt8(s []int8, from, to int) []int8 {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceInt16(s []int16, from, to int) []int16 {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceInt32(s []int32, from, to int) []int32 {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceInt64(s []int64, from, to int) []int64 {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceUint(s []uint, from, to int) []uint {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceUint8(s []uint8, from, to int) []uint8 {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceUint16(s []uint16, from, to int) []uint16 {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceUint32(s []uint32, from, to int) []uint32 {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceUint64(s []uint64, from, to int) []uint64 {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceUintptr(s []uintptr, from, to int) []uintptr {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceFloat32(s []float32, from, to int) []float32 {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceFloat64(s []float64, from, to int) []float64 {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceRune(s []rune, from, to int) []rune {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceByte(s []byte, from, to int) []byte {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceComplex64(s []complex64, from, to int) []complex64 {
	return append(s[:from], s[to:]...)
}

//删除子切片，子切片的位置是从from（包括）到to（不包括）
func DeleteSubSliceComplex128(s []complex128, from, to int) []complex128 {
	return append(s[:from], s[to:]...)
}
