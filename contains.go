// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//返回切片sli中是否包含单个元素v
func Contains(s []interface{}, v interface{}) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsString(s []string, v string) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsInt(s []int, v int) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsInt8(s []int8, v int8) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsInt16(s []int16, v int16) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsInt32(s []int32, v int32) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsInt64(s []int64, v int64) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUint(s []uint, v uint) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUint8(s []uint8, v uint8) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUint16(s []uint16, v uint16) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUint32(s []uint32, v uint32) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUint64(s []uint64, v uint64) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUintptr(s []uintptr, v uintptr) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsFloat32(s []float32, v float32) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsFloat64(s []float64, v float64) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsRune(s []rune, v rune) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsByte(s []byte, v byte) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsComplex64(s []complex64, v complex64) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsComplex128(s []complex128, v complex128) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsBool(s []bool, v bool) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}
