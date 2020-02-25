// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnore(sli []interface{}, v interface{}) []interface{} {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreString(sli []string, v string) []string {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreInt(sli []int, v int) []int {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreInt8(sli []int8, v int8) []int8 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreInt16(sli []int16, v int16) []int16 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreInt32(sli []int32, v int32) []int32 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreInt64(sli []int64, v int64) []int64 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUint(sli []uint, v uint) []uint {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUint8(sli []uint8, v uint8) []uint8 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUint16(sli []uint16, v uint16) []uint16 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUint32(sli []uint32, v uint32) []uint32 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUint64(sli []uint64, v uint64) []uint64 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreUintptr(sli []uintptr, v uintptr) []uintptr {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreFloat32(sli []float32, v float32) []float32 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreFloat64(sli []float64, v float64) []float64 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreRune(sli []rune, v rune) []rune {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreByte(sli []byte, v byte) []byte {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreBool(sli []bool, v bool) []bool {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreComplex64(sli []complex64, v complex64) []complex64 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}

//向切片Sli中插入没出现过的元素V，如果切片中有V，则不插入
func InsertIgnoreComplex128(sli []complex128, v complex128) []complex128 {
	for _, val := range sli {
		if val == v {
			return sli
		}
	}
	sli = append(sli, v)
	return sli
}
