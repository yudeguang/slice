// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//返回右往左最多n个元素
func Right(s []interface{}, n int) []interface{} {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightString(s []string, n int) []string {
	if len(s) <= n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightInt(s []int, n int) []int {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightInt8(s []int8, n int) []int8 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightInt16(s []int16, n int) []int16 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightInt32(s []int32, n int) []int32 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightInt64(s []int64, n int) []int64 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightUint(s []uint, n int) []uint {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightUint8(s []uint8, n int) []uint8 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightUint16(s []uint16, n int) []uint16 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightUint32(s []uint32, n int) []uint32 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightUint64(s []uint64, n int) []uint64 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightUintptr(s []uintptr, n int) []uintptr {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightFloat32(s []float32, n int) []float32 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightFloat64(s []float64, n int) []float64 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightRune(s []rune, n int) []rune {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightByte(s []byte, n int) []byte {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightComplex64(s []complex64, n int) []complex64 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightComplex128(s []complex128, n int) []complex128 {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}

//返回右往左最多n个元素
func RightBool(s []bool, n int) []bool {
	if len(s) < n {
		return s
	}
	if n <= 0 {
		return nil
	}
	return s[len(s)-n:]
}
