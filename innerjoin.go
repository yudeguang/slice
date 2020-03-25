// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

import (
	"github.com/yudeguang/hashset"
)

//返回去重复后两个切片的交集
func InnerJoin(m, n []interface{}) []interface{} {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinAny(hashset.New(m...), hashset.New(n...)).ToSlice()
	}

	distinctM, distinctN := Distinct(m), Distinct(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]interface{}, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if Contains(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func InnerJoinString(m, n []string) []string {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinString(hashset.NewString(m...), hashset.NewString(n...)).ToSlice()
	}
	distinctM := DistinctString(m)
	distinctN := DistinctString(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]string, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsString(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func InnerJoinInt(m, n []int) []int {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinInt(hashset.NewInt(m...), hashset.NewInt(n...)).ToSlice()
	}
	distinctM := DistinctInt(m)
	distinctN := DistinctInt(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]int, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsInt(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func InnerJoinInt8(m, n []int8) []int8 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinInt8(hashset.NewInt8(m...), hashset.NewInt8(n...)).ToSlice()
	}
	distinctM := DistinctInt8(m)
	distinctN := DistinctInt8(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]int8, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsInt8(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func InnerJoinInt16(m, n []int16) []int16 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinInt16(hashset.NewInt16(m...), hashset.NewInt16(n...)).ToSlice()
	}
	distinctM := DistinctInt16(m)
	distinctN := DistinctInt16(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]int16, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsInt16(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func InnerJoinInt32(m, n []int32) []int32 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinInt32(hashset.NewInt32(m...), hashset.NewInt32(n...)).ToSlice()
	}
	distinctM := DistinctInt32(m)
	distinctN := DistinctInt32(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]int32, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsInt32(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinInt64(m, n []int64) []int64 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinInt64(hashset.NewInt64(m...), hashset.NewInt64(n...)).ToSlice()
	}
	distinctM := DistinctInt64(m)
	distinctN := DistinctInt64(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]int64, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsInt64(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinUint(m, n []uint) []uint {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinUint(hashset.NewUint(m...), hashset.NewUint(n...)).ToSlice()
	}
	distinctM := DistinctUint(m)
	distinctN := DistinctUint(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]uint, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsUint(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinUint8(m, n []uint8) []uint8 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinUint8(hashset.NewUint8(m...), hashset.NewUint8(n...)).ToSlice()
	}
	distinctM := DistinctUint8(m)
	distinctN := DistinctUint8(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]uint8, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsUint8(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinUint16(m, n []uint16) []uint16 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinUint16(hashset.NewUint16(m...), hashset.NewUint16(n...)).ToSlice()
	}
	distinctM := DistinctUint16(m)
	distinctN := DistinctUint16(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]uint16, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsUint16(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinUint32(m, n []uint32) []uint32 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinUint32(hashset.NewUint32(m...), hashset.NewUint32(n...)).ToSlice()
	}
	distinctM := DistinctUint32(m)
	distinctN := DistinctUint32(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]uint32, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsUint32(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinUint64(m, n []uint64) []uint64 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinUint64(hashset.NewUint64(m...), hashset.NewUint64(n...)).ToSlice()
	}
	distinctM := DistinctUint64(m)
	distinctN := DistinctUint64(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]uint64, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsUint64(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinUintptr(m, n []uintptr) []uintptr {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinUintptr(hashset.NewUintptr(m...), hashset.NewUintptr(n...)).ToSlice()
	}
	distinctM := DistinctUintptr(m)
	distinctN := DistinctUintptr(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]uintptr, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsUintptr(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinFloat32(m, n []float32) []float32 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinFloat32(hashset.NewFloat32(m...), hashset.NewFloat32(n...)).ToSlice()
	}
	distinctM := DistinctFloat32(m)
	distinctN := DistinctFloat32(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]float32, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsFloat32(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinFloat64(m, n []float64) []float64 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinFloat64(hashset.NewFloat64(m...), hashset.NewFloat64(n...)).ToSlice()
	}
	distinctM := DistinctFloat64(m)
	distinctN := DistinctFloat64(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]float64, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsFloat64(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinRune(m, n []rune) []rune {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinRune(hashset.NewRune(m...), hashset.NewRune(n...)).ToSlice()
	}
	distinctM := DistinctRune(m)
	distinctN := DistinctRune(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]rune, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsRune(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinByte(m, n []byte) []byte {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinByte(hashset.NewByte(m...), hashset.NewByte(n...)).ToSlice()
	}
	distinctM := DistinctByte(m)
	distinctN := DistinctByte(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]byte, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsByte(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinComplex64(m, n []complex64) []complex64 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinComplex64(hashset.NewComplex64(m...), hashset.NewComplex64(n...)).ToSlice()
	}
	distinctM := DistinctComplex64(m)
	distinctN := DistinctComplex64(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]complex64, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsComplex64(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinComplex128(m, n []complex128) []complex128 {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	if L := len(m) * len(n); L > hashAlgorithmSwitchNum {
		return hashset.InnerJoinComplex128(hashset.NewComplex128(m...), hashset.NewComplex128(n...)).ToSlice()
	}
	distinctM := DistinctComplex128(m)
	distinctN := DistinctComplex128(n)
	lenM, lenN := len(distinctM), len(distinctN)
	result := make([]complex128, 0, minInt(lenM, lenN))
	for i := 0; i < lenM; i++ {
		if ContainsComplex128(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}

//返回去重复后两个切片的交集
func innerJoinBool(m, n []bool) []bool {
	if len(m) == 0 || len(n) == 0 {
		return nil
	}
	distinctM := DistinctBool(m)
	distinctN := DistinctBool(n)
	lenM := len(distinctM)
	result := make([]bool, 0, 2)
	for i := 0; i < lenM; i++ {
		if ContainsBool(distinctN, distinctM[i]) {
			result = append(result, distinctM[i])
		}
	}
	return result
}
