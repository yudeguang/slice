// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

func fmtCount(count ...int) int {
	if len(count) >= 1 {
		return count[0]
	}
	return -1
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func Remove(sli []interface{}, v interface{}, count ...int) []interface{} {
	n := fmtCount(count...)
	if len(sli) == 0 || v == nil || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]interface{}, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []interface{}{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveString(sli []string, v string, count ...int) []string {
	n := fmtCount(count...)
	if len(sli) == 0 || v == "" || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]string, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []string{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveInt(sli []int, v int, count ...int) []int {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]int, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []int{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveInt8(sli []int8, v int8, count ...int) []int8 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]int8, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []int8{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveInt16(sli []int16, v int16, count ...int) []int16 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]int16, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []int16{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveInt32(sli []int32, v int32, count ...int) []int32 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]int32, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []int32{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveInt64(sli []int64, v int64, count ...int) []int64 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]int64, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []int64{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveUint(sli []uint, v uint, count ...int) []uint {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]uint, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []uint{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveUint8(sli []uint8, v uint8, count ...int) []uint8 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]uint8, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []uint8{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveUint16(sli []uint16, v uint16, count ...int) []uint16 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]uint16, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []uint16{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveUint32(sli []uint32, v uint32, count ...int) []uint32 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]uint32, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []uint32{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveUint64(sli []uint64, v uint64, count ...int) []uint64 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]uint64, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []uint64{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveUintptr(sli []uintptr, v uintptr, count ...int) []uintptr {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]uintptr, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []uintptr{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveFloat32(sli []float32, v float32, count ...int) []float32 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]float32, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []float32{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveFloat64(sli []float64, v float64, count ...int) []float64 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]float64, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []float64{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveRune(sli []rune, v rune, count ...int) []rune {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]rune, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []rune{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveByte(sli []byte, v byte, count ...int) []byte {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]byte, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []byte{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveBool(sli []bool, v bool, count ...int) []bool {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]bool, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []bool{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveComplex64(sli []complex64, v complex64, count ...int) []complex64 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]complex64, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []complex64{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将sli中前n个子元素v删除后的切片，如果这填，或者n<0会删除所有子元素v
func RemoveComplex128(sli []complex128, v complex128, count ...int) []complex128 {
	n := fmtCount(count...)
	if len(sli) == 0 || n == 0 {
		return sli
	}
	//全部删除
	if n < 0 {
		ret := make([]complex128, 0, len(sli)-n)
		for _, val := range sli {
			if val != v {
				ret = append(ret, val)
			}
		}
		return ret
	}
	//删除n次
	ret := []complex128{}
	t := 0
	for i, val := range sli {
		if val != v {
			ret = append(ret, val)
		} else {
			t++
			if t == n {
				ret = append(ret, sli[i+1:]...)
				break
			}
		}
	}
	return ret
}
