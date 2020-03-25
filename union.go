// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func Union(m, n []interface{}) []interface{} {
	return Distinct(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionString(m, n []string) []string {
	return DistinctString(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionInt(m, n []int) []int {
	return DistinctInt(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionInt8(m, n []int8) []int8 {
	return DistinctInt8(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionInt16(m, n []int16) []int16 {
	return DistinctInt16(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionInt32(m, n []int32) []int32 {
	return DistinctInt32(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionInt64(m, n []int64) []int64 {
	return DistinctInt64(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionUint(m, n []uint) []uint {
	return DistinctUint(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionUint8(m, n []uint8) []uint8 {
	return DistinctUint8(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionUint16(m, n []uint16) []uint16 {
	return DistinctUint16(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionUint32(m, n []uint32) []uint32 {
	return DistinctUint32(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionUint64(m, n []uint64) []uint64 {
	return DistinctUint64(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionUintptr(m, n []uintptr) []uintptr {
	return DistinctUintptr(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionFloat32(m, n []float32) []float32 {
	return DistinctFloat32(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionFloat64(m, n []float64) []float64 {
	return DistinctFloat64(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionRune(m, n []rune) []rune {
	return DistinctRune(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionByte(m, n []byte) []byte {
	return DistinctByte(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionComplex64(m, n []complex64) []complex64 {
	return DistinctComplex64(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionComplex128(m, n []complex128) []complex128 {
	return DistinctComplex128(append(append(m[:0:0], m...), n...))
}

//获得两个切片去重后的合集,也即m+n，注意union 与 union all的区别，这里并不实现union all
func UnionBool(m, n []bool) []bool {
	return DistinctBool(append(append(m[:0:0], m...), n...))
}
