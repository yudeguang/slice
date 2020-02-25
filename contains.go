// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//Contains无需启用set，因为初始化成本已经高于遍历一次的成本
//返回切片sli中是否包含单个元素v
func Contains(sli []interface{}, v interface{}) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsString(sli []string, v string) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsInt(sli []int, v int) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsInt8(sli []int8, v int8) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsInt16(sli []int16, v int16) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsInt32(sli []int32, v int32) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsInt64(sli []int64, v int64) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUint(sli []uint, v uint) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUint8(sli []uint8, v uint8) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUint16(sli []uint16, v uint16) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUint32(sli []uint32, v uint32) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUint64(sli []uint64, v uint64) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsUintptr(sli []uintptr, v uintptr) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsFloat32(sli []float32, v float32) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsFloat64(sli []float64, v float64) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsRune(sli []rune, v rune) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsByte(sli []byte, v byte) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsComplex64(sli []complex64, v complex64) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}

//返回切片sli中是否包含单个元素v
func ContainsComplex128(sli []complex128, v complex128) bool {
	for _, vv := range sli {
		if vv == v {
			return true
		}
	}
	return false
}
