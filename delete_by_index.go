// DeleteByIndexright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a DeleteByIndex of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

/*
// 第一种方法（保持剩余元素的次序）：
s = append(s[:i], s[i+1:]...)

// 第二种方法（保持剩余元素的次序）：
s = s[:i + copy(s[i:], s[i+1:])]

// 上面两种方法都需要复制len(s)-i-1个元素。

// 第三种方法（不保持剩余元素的次序）：
s[i] = s[len(s)-1]
s = s[:len(s)-1]

//在此，只实现第一种方法
*/
//删除切片中第i个元素
func DeleteByIndex(s []interface{}, i int) []interface{} {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexString(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexInt(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexInt8(s []int8, i int) []int8 {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexInt16(s []int16, i int) []int16 {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexInt32(s []int32, i int) []int32 {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexInt64(s []int64, i int) []int64 {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUint(s []uint, i int) []uint {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUint8(s []uint8, i int) []uint8 {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUint16(s []uint16, i int) []uint16 {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUint32(s []uint32, i int) []uint32 {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUint64(s []uint64, i int) []uint64 {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUintptr(s []uintptr, i int) []uintptr {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexFloat32(s []float32, i int) []float32 {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexFloat64(s []float64, i int) []float64 {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexRune(s []rune, i int) []rune {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexByte(s []byte, i int) []byte {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexComplex64(s []complex64, i int) []complex64 {
	return append(s[:i], s[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexComplex128(s []complex128, i int) []complex128 {
	return append(s[:i], s[i+1:]...)
}
