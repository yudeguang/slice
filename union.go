// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func Union(m, n []interface{}) []interface{} {
	result := make([]interface{}, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return Distinct(result)
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionString(m, n []string) []string {
	result := make([]string, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctString(result)
}

//返回去重复后两个切片的并集
func UnionInt(m, n []int) []int {
	result := make([]int, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctInt(result)
}

//返回去重复后两个切片的并集
func UnionInt8(m, n []int8) []int8 {
	result := make([]int8, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctInt8(result)
}

//返回去重复后两个切片的并集
func UnionInt16(m, n []int16) []int16 {
	result := make([]int16, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctInt16(result)
}

//返回去重复后两个切片的并集
func UnionInt32(m, n []int32) []int32 {
	result := make([]int32, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctInt32(result)
}

//返回去重复后两个切片的并集
func UnionInt64(m, n []int64) []int64 {
	result := make([]int64, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctInt64(result)
}

//返回去重复后两个切片的并集
func UnionUint(m, n []uint) []uint {
	result := make([]uint, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctUint(result)
}

//返回去重复后两个切片的并集
func UnionUint8(m, n []uint8) []uint8 {
	result := make([]uint8, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctUint8(result)
}

//返回去重复后两个切片的并集
func UnionUint16(m, n []uint16) []uint16 {
	result := make([]uint16, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctUint16(result)
}

//返回去重复后两个切片的并集
func UnionUint32(m, n []uint32) []uint32 {
	result := make([]uint32, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctUint32(result)
}

//返回去重复后两个切片的并集
func UnionUint64(m, n []uint64) []uint64 {
	result := make([]uint64, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctUint64(result)
}

//返回去重复后两个切片的并集
func UnionUintptr(m, n []uintptr) []uintptr {
	result := make([]uintptr, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctUintptr(result)
}

//返回去重复后两个切片的并集
func UnionFloat32(m, n []float32) []float32 {
	result := make([]float32, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctFloat32(result)
}

//返回去重复后两个切片的并集
func UnionFloat64(m, n []float64) []float64 {
	result := make([]float64, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctFloat64(result)
}

//返回去重复后两个切片的并集
func UnionRune(m, n []rune) []rune {
	result := make([]rune, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctRune(result)
}

//返回去重复后两个切片的并集
func UnionByte(m, n []byte) []byte {
	result := make([]byte, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctByte(result)
}

//返回去重复后两个切片的并集
func UnionComplex64(m, n []complex64) []complex64 {
	result := make([]complex64, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctComplex64(result)
}

//返回去重复后两个切片的并集
func UnionComplex128(m, n []complex128) []complex128 {
	result := make([]complex128, 0, len(m)+len(n))
	copy(result, m)
	copy(result[len(m):], n)
	//Distinct函数会自动判断是否足够长，会根据程度决定是否启用hashSet算法
	return DistinctComplex128(result)
}
