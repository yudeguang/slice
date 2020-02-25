// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//返回集合A不在集合B中的所有值，不做去重操作
func NotIn(a, b []interface{}) []interface{} {
	result := make([]interface{}, 0, len(a))
	for _, v := range a {
		if !Contains(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回集合A不在集合B中的所有值，不做去重操作
func NotInString(a, b []string) []string {
	result := make([]string, 0, len(a))
	for _, v := range a {
		if !ContainsString(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInInt(a, b []int) []int {
	result := make([]int, 0, len(a))
	for _, v := range a {
		if !ContainsInt(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInInt8(a, b []int8) []int8 {
	result := make([]int8, 0, len(a))
	for _, v := range a {
		if !ContainsInt8(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInInt16(a, b []int16) []int16 {
	result := make([]int16, 0, len(a))
	for _, v := range a {
		if !ContainsInt16(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInInt32(a, b []int32) []int32 {
	result := make([]int32, 0, len(a))
	for _, v := range a {
		if !ContainsInt32(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInInt64(a, b []int64) []int64 {
	result := make([]int64, 0, len(a))
	for _, v := range a {
		if !ContainsInt64(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInUint(a, b []uint) []uint {
	result := make([]uint, 0, len(a))
	for _, v := range a {
		if !ContainsUint(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInUint8(a, b []uint8) []uint8 {
	result := make([]uint8, 0, len(a))
	for _, v := range a {
		if !ContainsUint8(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInUint16(a, b []uint16) []uint16 {
	result := make([]uint16, 0, len(a))
	for _, v := range a {
		if !ContainsUint16(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInUint32(a, b []uint32) []uint32 {
	result := make([]uint32, 0, len(a))
	for _, v := range a {
		if !ContainsUint32(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInUint64(a, b []uint64) []uint64 {
	result := make([]uint64, 0, len(a))
	for _, v := range a {
		if !ContainsUint64(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInUintptr(a, b []uintptr) []uintptr {
	result := make([]uintptr, 0, len(a))
	for _, v := range a {
		if !ContainsUintptr(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInFloat32(a, b []float32) []float32 {
	result := make([]float32, 0, len(a))
	for _, v := range a {
		if !ContainsFloat32(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInFloat64(a, b []float64) []float64 {
	result := make([]float64, 0, len(a))
	for _, v := range a {
		if !ContainsFloat64(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInRune(a, b []rune) []rune {
	result := make([]rune, 0, len(a))
	for _, v := range a {
		if !ContainsRune(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInByte(a, b []byte) []byte {
	result := make([]byte, 0, len(a))
	for _, v := range a {
		if !ContainsByte(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInComplex64(a, b []complex64) []complex64 {
	result := make([]complex64, 0, len(a))
	for _, v := range a {
		if !ContainsComplex64(b, v) {
			result = append(result, v)
		}
	}
	return result
}

//返回去重复后两个切片的交集
func NotInComplex128(a, b []complex128) []complex128 {
	result := make([]complex128, 0, len(a))
	for _, v := range a {
		if !ContainsComplex128(b, v) {
			result = append(result, v)
		}
	}
	return result
}
