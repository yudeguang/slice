// Reverseright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//对切片进行倒序排列
func Reverse(s []interface{}) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseString(s []string) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseInt(s []int) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseInt8(s []int8) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseInt16(s []int16) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseInt32(s []int32) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseInt64(s []int64) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseUint(s []uint) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseUint8(s []uint8) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseUint16(s []uint16) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseUint32(s []uint32) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseUint64(s []uint64) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseUintptr(s []uintptr) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseFloat32(s []float32) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseFloat64(s []float64) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseRune(s []rune) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseByte(s []byte) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseComplex64(s []complex64) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseComplex128(s []complex128) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

//对切片进行倒序排列
func ReverseBool(s []bool) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}
