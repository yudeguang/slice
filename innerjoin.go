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
//返回去重复后两个切片的交集
func InnerJoin(a, b []interface{}) []interface{} {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.New(a...)
		setB := hashset.New(b...)
		set := hashset.InnerJoinAny(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA, distinctB := Distinct(a), Distinct(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]interface{}, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if Contains(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}

}

//返回去重复后两个切片的交集
func InnerJoinString(a, b []string) []string {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewString(a...)
		setB := hashset.NewString(b...)
		set := hashset.InnerJoinString(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctString(a)
		distinctB := DistinctString(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]string, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsString(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func InnerJoinInt(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewInt(a...)
		setB := hashset.NewInt(b...)
		set := hashset.InnerJoinInt(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctInt(a)
		distinctB := DistinctInt(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]int, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsInt(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func InnerJoinInt8(a, b []int8) []int8 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewInt8(a...)
		setB := hashset.NewInt8(b...)
		set := hashset.InnerJoinInt8(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctInt8(a)
		distinctB := DistinctInt8(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]int8, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsInt8(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func InnerJoinInt16(a, b []int16) []int16 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewInt16(a...)
		setB := hashset.NewInt16(b...)
		set := hashset.InnerJoinInt16(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctInt16(a)
		distinctB := DistinctInt16(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]int16, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsInt16(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func InnerJoinInt32(a, b []int32) []int32 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewInt32(a...)
		setB := hashset.NewInt32(b...)
		set := hashset.InnerJoinInt32(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctInt32(a)
		distinctB := DistinctInt32(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]int32, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsInt32(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinInt64(a, b []int64) []int64 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewInt64(a...)
		setB := hashset.NewInt64(b...)
		set := hashset.InnerJoinInt64(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctInt64(a)
		distinctB := DistinctInt64(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]int64, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsInt64(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinUint(a, b []uint) []uint {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewUint(a...)
		setB := hashset.NewUint(b...)
		set := hashset.InnerJoinUint(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctUint(a)
		distinctB := DistinctUint(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]uint, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsUint(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinUint8(a, b []uint8) []uint8 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewUint8(a...)
		setB := hashset.NewUint8(b...)
		set := hashset.InnerJoinUint8(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctUint8(a)
		distinctB := DistinctUint8(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]uint8, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsUint8(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinUint16(a, b []uint16) []uint16 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewUint16(a...)
		setB := hashset.NewUint16(b...)
		set := hashset.InnerJoinUint16(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctUint16(a)
		distinctB := DistinctUint16(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]uint16, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsUint16(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinUint32(a, b []uint32) []uint32 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewUint32(a...)
		setB := hashset.NewUint32(b...)
		set := hashset.InnerJoinUint32(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctUint32(a)
		distinctB := DistinctUint32(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]uint32, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsUint32(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinUint64(a, b []uint64) []uint64 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewUint64(a...)
		setB := hashset.NewUint64(b...)
		set := hashset.InnerJoinUint64(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctUint64(a)
		distinctB := DistinctUint64(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]uint64, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsUint64(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinUintptr(a, b []uintptr) []uintptr {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewUintptr(a...)
		setB := hashset.NewUintptr(b...)
		set := hashset.InnerJoinUintptr(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctUintptr(a)
		distinctB := DistinctUintptr(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]uintptr, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsUintptr(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinFloat32(a, b []float32) []float32 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewFloat32(a...)
		setB := hashset.NewFloat32(b...)
		set := hashset.InnerJoinFloat32(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctFloat32(a)
		distinctB := DistinctFloat32(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]float32, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsFloat32(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinFloat64(a, b []float64) []float64 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewFloat64(a...)
		setB := hashset.NewFloat64(b...)
		set := hashset.InnerJoinFloat64(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctFloat64(a)
		distinctB := DistinctFloat64(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]float64, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsFloat64(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinRune(a, b []rune) []rune {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewRune(a...)
		setB := hashset.NewRune(b...)
		set := hashset.InnerJoinRune(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctRune(a)
		distinctB := DistinctRune(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]rune, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsRune(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinByte(a, b []byte) []byte {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewByte(a...)
		setB := hashset.NewByte(b...)
		set := hashset.InnerJoinByte(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctByte(a)
		distinctB := DistinctByte(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]byte, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsByte(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinComplex64(a, b []complex64) []complex64 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewComplex64(a...)
		setB := hashset.NewComplex64(b...)
		set := hashset.InnerJoinComplex64(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctComplex64(a)
		distinctB := DistinctComplex64(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]complex64, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsComplex64(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}

//返回去重复后两个切片的交集
func innerJoinComplex128(a, b []complex128) []complex128 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if L := len(a) * len(b); L > bestNum {
		setA := hashset.NewComplex128(a...)
		setB := hashset.NewComplex128(b...)
		set := hashset.InnerJoinComplex128(setA, setB)
		result := set.ToSlice()
		setA.Clear()
		setB.Clear()
		set.Clear()
		return result
	} else {
		distinctA := DistinctComplex128(a)
		distinctB := DistinctComplex128(b)
		lenA, lenB := len(distinctA), len(distinctB)
		result := make([]complex128, 0, maxInt(lenA, lenB))
		for i := 0; i < lenA; i++ {
			if ContainsComplex128(distinctB, distinctA[i]) {
				result = append(result, distinctA[i])
			}
		}
		return result
	}
}
