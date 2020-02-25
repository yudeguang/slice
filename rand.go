// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

import (
	"math/rand"
	"time"
)

//返回随机打乱后的切片
func Rand(s []interface{}) []interface{} {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]interface{}, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回随机打乱后的切片
func RandString(s []string) []string {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]string, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandInt(s []int) []int {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]int, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandInt8(s []int8) []int8 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]int8, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandInt16(s []int16) []int16 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]int16, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandInt32(s []int32) []int32 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]int32, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandInt64(s []int64) []int64 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]int64, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandUint(s []uint) []uint {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]uint, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandUint8(s []uint8) []uint8 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]uint8, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandUint16(s []uint16) []uint16 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]uint16, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandUint32(s []uint32) []uint32 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]uint32, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandUint64(s []uint64) []uint64 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]uint64, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandUintptr(s []uintptr) []uintptr {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]uintptr, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandFloat32(s []float32) []float32 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]float32, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandFloat64(s []float64) []float64 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]float64, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandRune(s []rune) []rune {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]rune, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandByte(s []byte) []byte {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]byte, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandBool(s []bool) []bool {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]bool, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandComplex64(s []complex64) []complex64 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]complex64, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}

//返回切片s中是否包含单个元素v
func RandComplex128(s []complex128) []complex128 {
	rand.Seed(time.Now().UnixNano())
	L := len(s)
	randedId := rand.Perm(L)
	randedsce := make([]complex128, 0, L)
	for _, i := range randedId {
		randedsce = append(randedsce, s[i])
	}
	return randedsce
}
