// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnore(s []interface{}, v interface{}) []interface{} {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreString(s []string, v string) []string {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreInt(s []int, v int) []int {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreInt8(s []int8, v int8) []int8 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreInt16(s []int16, v int16) []int16 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreInt32(s []int32, v int32) []int32 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreInt64(s []int64, v int64) []int64 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUint(s []uint, v uint) []uint {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUint8(s []uint8, v uint8) []uint8 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUint16(s []uint16, v uint16) []uint16 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUint32(s []uint32, v uint32) []uint32 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUint64(s []uint64, v uint64) []uint64 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUintptr(s []uintptr, v uintptr) []uintptr {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreFloat32(s []float32, v float32) []float32 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreFloat64(s []float64, v float64) []float64 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreRune(s []rune, v rune) []rune {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreByte(s []byte, v byte) []byte {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreBool(s []bool, v bool) []bool {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreComplex64(s []complex64, v complex64) []complex64 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreComplex128(s []complex128, v complex128) []complex128 {
	for i := range s {
		if s[i] == v {
			return s
		}
	}
	return append(s, v)
}
