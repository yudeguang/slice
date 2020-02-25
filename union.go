// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

import (
	"github.com/yudeguang/hashset"
)

//当数量较大时，一般大于1024左右set会有更高的效率
//获得两个切片去重后的合集
func Union(a, b []interface{}) []interface{} {

	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetAny{Items: make(map[interface{}]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := Distinct(a), Distinct(b)
		for _, v := range distinctA {
			if !Contains(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}

}

//获得两个切片去重后的合集
func UnionString(a, b []string) []string {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetString{Items: make(map[string]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctString(a), DistinctString(b)
		for _, v := range distinctA {
			if !ContainsString(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionInt(a, b []int) []int {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetInt{Items: make(map[int]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctInt(a), DistinctInt(b)
		for _, v := range distinctA {
			if !ContainsInt(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionInt8(a, b []int8) []int8 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetInt8{Items: make(map[int8]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctInt8(a), DistinctInt8(b)
		for _, v := range distinctA {
			if !ContainsInt8(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionInt16(a, b []int16) []int16 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetInt16{Items: make(map[int16]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctInt16(a), DistinctInt16(b)
		for _, v := range distinctA {
			if !ContainsInt16(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionInt32(a, b []int32) []int32 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetInt32{Items: make(map[int32]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctInt32(a), DistinctInt32(b)
		for _, v := range distinctA {
			if !ContainsInt32(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionInt64(a, b []int64) []int64 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetInt64{Items: make(map[int64]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctInt64(a), DistinctInt64(b)
		for _, v := range distinctA {
			if !ContainsInt64(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionUint(a, b []uint) []uint {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetUint{Items: make(map[uint]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctUint(a), DistinctUint(b)
		for _, v := range distinctA {
			if !ContainsUint(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionUint8(a, b []uint8) []uint8 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetUint8{Items: make(map[uint8]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctUint8(a), DistinctUint8(b)
		for _, v := range distinctA {
			if !ContainsUint8(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionUint16(a, b []uint16) []uint16 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetUint16{Items: make(map[uint16]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctUint16(a), DistinctUint16(b)
		for _, v := range distinctA {
			if !ContainsUint16(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionUint32(a, b []uint32) []uint32 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetUint32{Items: make(map[uint32]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctUint32(a), DistinctUint32(b)
		for _, v := range distinctA {
			if !ContainsUint32(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionUint64(a, b []uint64) []uint64 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetUint64{Items: make(map[uint64]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctUint64(a), DistinctUint64(b)
		for _, v := range distinctA {
			if !ContainsUint64(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionUintptr(a, b []uintptr) []uintptr {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetUintptr{Items: make(map[uintptr]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctUintptr(a), DistinctUintptr(b)
		for _, v := range distinctA {
			if !ContainsUintptr(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionFloat32(a, b []float32) []float32 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetFloat32{Items: make(map[float32]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctFloat32(a), DistinctFloat32(b)
		for _, v := range distinctA {
			if !ContainsFloat32(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionFloat64(a, b []float64) []float64 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetFloat64{Items: make(map[float64]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctFloat64(a), DistinctFloat64(b)
		for _, v := range distinctA {
			if !ContainsFloat64(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionRune(a, b []rune) []rune {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetRune{Items: make(map[rune]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctRune(a), DistinctRune(b)
		for _, v := range distinctA {
			if !ContainsRune(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionByte(a, b []byte) []byte {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetByte{Items: make(map[byte]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctByte(a), DistinctByte(b)
		for _, v := range distinctA {
			if !ContainsByte(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionComplex64(a, b []complex64) []complex64 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetComplex64{Items: make(map[complex64]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctComplex64(a), DistinctComplex64(b)
		for _, v := range distinctA {
			if !ContainsComplex64(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}

//返回去重复后两个切片的交集
func UnionComplex128(a, b []complex128) []complex128 {
	if L := len(a) * len(b); L > bestNum {
		set := &hashset.SetComplex128{Items: make(map[complex128]struct{}, L)}
		set.Add(a...)
		set.Add(b...)
		result := set.ToSlice()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := DistinctComplex128(a), DistinctComplex128(b)
		for _, v := range distinctA {
			if !ContainsComplex128(distinctB, v) {
				distinctB = append(distinctB, v)
			}
		}
		return distinctB
	}
}
