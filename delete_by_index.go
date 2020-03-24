// DeleteByIndexright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a DeleteByIndex of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

/*
// 第一种方法（保持剩余元素的次序）：
s = append(copyS[:i], copyS[i+1:]...)

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
	copyS := Copy(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexString(s []string, i int) []string {
	copyS := CopyString(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexInt(s []int, i int) []int {
	copyS := CopyInt(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexInt8(s []int8, i int) []int8 {
	copyS := CopyInt8(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexInt16(s []int16, i int) []int16 {
	copyS := CopyInt16(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexInt32(s []int32, i int) []int32 {
	copyS := CopyInt32(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexInt64(s []int64, i int) []int64 {
	copyS := CopyInt64(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUint(s []uint, i int) []uint {
	copyS := CopyUint(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUint8(s []uint8, i int) []uint8 {
	copyS := CopyUint8(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUint16(s []uint16, i int) []uint16 {
	copyS := CopyUint16(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUint32(s []uint32, i int) []uint32 {
	copyS := CopyUint32(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUint64(s []uint64, i int) []uint64 {
	copyS := CopyUint64(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexUintptr(s []uintptr, i int) []uintptr {
	copyS := CopyUintptr(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexFloat32(s []float32, i int) []float32 {
	copyS := CopyFloat32(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexFloat64(s []float64, i int) []float64 {
	copyS := CopyFloat64(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexRune(s []rune, i int) []rune {
	copyS := CopyRune(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexByte(s []byte, i int) []byte {
	copyS := CopyByte(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexComplex64(s []complex64, i int) []complex64 {
	copyS := CopyComplex64(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexComplex128(s []complex128, i int) []complex128 {
	copyS := CopyComplex128(s)
	return append(copyS[:i], copyS[i+1:]...)
}

//删除切片中第i个元素
func DeleteByIndexBool(s []bool, i int) []bool {
	copyS := CopyBool(s)
	return append(copyS[:i], copyS[i+1:]...)
}
