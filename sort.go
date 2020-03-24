// Sortright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

import (
	"sort"
)

// func Sort(s []interface{}) []interface{} {
// 	sort.Slice(s, func(i, j int) bool {
// 		return s[i] < s[j]
// 	})
// 	return s
// }
//切片排序
func SortString(s []string) {
	sort.Strings(s)
}

//切片排序
func SortInt(s []int) {
	sort.Ints(s)
}

//切片排序
func SortInt8(s []int8) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortInt16(s []int16) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortInt32(s []int32) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortInt64(s []int64) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortUint(s []uint) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortUint8(s []uint8) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortUint16(s []uint16) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortUint32(s []uint32) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortUint64(s []uint64) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortUintptr(s []uintptr) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortFloat32(s []float32) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortFloat64(s []float64) {
	sort.Float64s(s)
}

//切片排序
func SortRune(s []rune) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

//切片排序
func SortByte(s []byte) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// //切片排序
// func SortComplex64(s []complex64) []complex64 {
// 	sort.Slice(s, func(i, j int) bool {
// 		return s[i] < s[j]
// 	})
// 	return s
// }

// //切片排序
// func SortComplex128(s []complex128) []complex128 {
// 	sort.Slice(s, func(i, j int) bool {
// 		return s[i] < s[j]
// 	})
// 	return s
// }

////切片排序
// func SortBool(s []bool) {
// 	sort.Slice(s, func(i, j int) bool {
// 		return s[i] < s[j]
// 	})
// }
