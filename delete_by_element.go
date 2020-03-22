// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

func fmtNum(num ...int) int {
	if len(num) >= 1 {
		return num[0]
	}
	return -1
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElement(s []interface{}, element interface{}, num ...int) []interface{} {
	n := fmtNum(num...)
	if len(s) == 0 || element == nil || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]interface{}, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := make([]interface{}, 0, len(s))
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementString(s []string, element string, num ...int) []string {
	n := fmtNum(num...)
	if len(s) == 0 || element == "" || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]string, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []string{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementInt(s []int, element int, num ...int) []int {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]int, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []int{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementInt8(s []int8, element int8, num ...int) []int8 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]int8, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []int8{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementInt16(s []int16, element int16, num ...int) []int16 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]int16, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []int16{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementInt32(s []int32, element int32, num ...int) []int32 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]int32, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []int32{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementInt64(s []int64, element int64, num ...int) []int64 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]int64, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []int64{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementUint(s []uint, element uint, num ...int) []uint {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]uint, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []uint{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementUint8(s []uint8, element uint8, num ...int) []uint8 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]uint8, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []uint8{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementUint16(s []uint16, element uint16, num ...int) []uint16 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]uint16, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []uint16{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementUint32(s []uint32, element uint32, num ...int) []uint32 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]uint32, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []uint32{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementUint64(s []uint64, element uint64, num ...int) []uint64 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]uint64, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []uint64{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementUintptr(s []uintptr, element uintptr, num ...int) []uintptr {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]uintptr, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []uintptr{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementFloat32(s []float32, element float32, num ...int) []float32 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]float32, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []float32{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementFloat64(s []float64, element float64, num ...int) []float64 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]float64, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []float64{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementRune(s []rune, element rune, num ...int) []rune {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]rune, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []rune{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementByte(s []byte, element byte, num ...int) []byte {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]byte, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []byte{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementBool(s []bool, element bool, num ...int) []bool {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]bool, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []bool{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementComplex64(s []complex64, element complex64, num ...int) []complex64 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]complex64, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []complex64{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}

//返回将s中删除前n个元素e后的切片，如果num不填，或者num<0会删除所有子元素e
func DeleteByElementComplex128(s []complex128, element complex128, num ...int) []complex128 {
	n := fmtNum(num...)
	if len(s) == 0 || n == 0 {
		return s
	}
	//全部删除
	if n < 0 {
		ret := make([]complex128, 0, len(s))
		for i := range s {
			if s[i] != element {
				ret = append(ret, s[i])
			}
		}
		return ret
	}
	//删除n次
	ret := []complex128{}
	t := 0
	for i := range s {
		if s[i] != element {
			ret = append(ret, s[i])
		} else {
			t++
			if t == n {
				ret = append(ret, s[i+1:]...)
				break
			}
		}
	}
	return ret
}
