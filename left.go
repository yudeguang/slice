// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//返回从左往右最多n个元素
func Left(s []interface{}, n int) []interface{} {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftString(s []string, n int) []string {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftInt(s []int, n int) []int {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftInt8(s []int8, n int) []int8 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftInt16(s []int16, n int) []int16 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftInt32(s []int32, n int) []int32 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftInt64(s []int64, n int) []int64 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftUint(s []uint, n int) []uint {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftUint8(s []uint8, n int) []uint8 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftUint16(s []uint16, n int) []uint16 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftUint32(s []uint32, n int) []uint32 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftUint64(s []uint64, n int) []uint64 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftUintptr(s []uintptr, n int) []uintptr {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftFloat32(s []float32, n int) []float32 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftFloat64(s []float64, n int) []float64 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftRune(s []rune, n int) []rune {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftByte(s []byte, n int) []byte {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftComplex64(s []complex64, n int) []complex64 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftComplex128(s []complex128, n int) []complex128 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}

//返回从左往右最多n个元素
func LeftBool(s []bool, n int) []bool {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[:n]
}
