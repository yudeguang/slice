// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

import "github.com/yudeguang/hashset"

//当数量较大时，一般大于1024左右set会有更高的效率
//返回去重复后的数据
func Distinct(s []interface{}) []interface{} {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.New(s...).ToSlice()
	}
	result := make([]interface{}, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !Contains(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result

}

//返回去重复后的数据
func DistinctString(s []string) []string {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewString(s...).ToSlice()
	}
	result := make([]string, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsString(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctInt(s []int) []int {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewInt(s...).ToSlice()
	}
	result := make([]int, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsInt(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctInt8(s []int8) []int8 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewInt8(s...).ToSlice()
	}
	result := make([]int8, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsInt8(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctInt16(s []int16) []int16 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewInt16(s...).ToSlice()
	}
	result := make([]int16, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsInt16(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctInt32(s []int32) []int32 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewInt32(s...).ToSlice()
	}
	result := make([]int32, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsInt32(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctInt64(s []int64) []int64 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewInt64(s...).ToSlice()
	}
	result := make([]int64, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsInt64(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctUint(s []uint) []uint {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewUint(s...).ToSlice()
	}
	result := make([]uint, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsUint(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctUint8(s []uint8) []uint8 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewUint8(s...).ToSlice()
	}
	result := make([]uint8, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsUint8(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctUint16(s []uint16) []uint16 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewUint16(s...).ToSlice()
	}
	result := make([]uint16, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsUint16(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctUint32(s []uint32) []uint32 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewUint32(s...).ToSlice()
	}
	result := make([]uint32, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsUint32(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctUint64(s []uint64) []uint64 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewUint64(s...).ToSlice()
	}
	result := make([]uint64, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsUint64(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctUintptr(s []uintptr) []uintptr {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewUintptr(s...).ToSlice()
	}
	result := make([]uintptr, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsUintptr(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctFloat32(s []float32) []float32 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewFloat32(s...).ToSlice()
	}
	result := make([]float32, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsFloat32(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctFloat64(s []float64) []float64 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewFloat64(s...).ToSlice()
	}
	result := make([]float64, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsFloat64(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctRune(s []rune) []rune {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewRune(s...).ToSlice()
	}
	result := make([]rune, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsRune(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctByte(s []byte) []byte {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewByte(s...).ToSlice()
	}
	result := make([]byte, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsByte(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctComplex64(s []complex64) []complex64 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewComplex64(s...).ToSlice()
	}
	result := make([]complex64, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsComplex64(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

//返回去重复后的数据
func DistinctComplex128(s []complex128) []complex128 {
	L := len(s)
	if L == 0 {
		return nil
	}
	if L > HashAlgorithmSwitchNum {
		return hashset.NewComplex128(s...).ToSlice()
	}
	result := make([]complex128, 0, L)
	result = append(result, s[0])
	for i := 1; i < L; i++ {
		if !ContainsComplex128(result, s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}
