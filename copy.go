// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//安全的复制切片
func Copy(s []interface{}) []interface{} {
	if s == nil {
		return nil
	}
	c := make([]interface{}, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyString(s []string) []string {
	if s == nil {
		return nil
	}
	c := make([]string, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyInt(s []int) []int {
	if s == nil {
		return nil
	}
	c := make([]int, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyInt8(s []int8) []int8 {
	if s == nil {
		return nil
	}
	c := make([]int8, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyInt16(s []int16) []int16 {
	if s == nil {
		return nil
	}
	c := make([]int16, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyInt32(s []int32) []int32 {
	if s == nil {
		return nil
	}
	c := make([]int32, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyInt64(s []int64) []int64 {
	if s == nil {
		return nil
	}
	c := make([]int64, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyUint(s []uint) []uint {
	if s == nil {
		return nil
	}
	c := make([]uint, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyUint8(s []uint8) []uint8 {
	if s == nil {
		return nil
	}
	c := make([]uint8, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyUint16(s []uint16) []uint16 {
	if s == nil {
		return nil
	}
	c := make([]uint16, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyUint32(s []uint32) []uint32 {
	if s == nil {
		return nil
	}
	c := make([]uint32, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyUint64(s []uint64) []uint64 {
	if s == nil {
		return nil
	}
	c := make([]uint64, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyUintptr(s []uintptr) []uintptr {
	if s == nil {
		return nil
	}
	c := make([]uintptr, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyFloat32(s []float32) []float32 {
	if s == nil {
		return nil
	}
	c := make([]float32, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyFloat64(s []float64) []float64 {
	if s == nil {
		return nil
	}
	c := make([]float64, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyRune(s []rune) []rune {
	if s == nil {
		return nil
	}
	c := make([]rune, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyByte(s []byte) []byte {
	if s == nil {
		return nil
	}
	c := make([]byte, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyComplex64(s []complex64) []complex64 {
	if s == nil {
		return nil
	}
	c := make([]complex64, len(s))
	copy(c, s)
	return c
}

//安全的复制切片
func CopyComplex128(s []complex128) []complex128 {
	if s == nil {
		return nil
	}
	c := make([]complex128, len(s))
	copy(c, s)
	return c
}
