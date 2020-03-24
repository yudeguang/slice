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

/*
https://github.com/golang/go/wiki/SliceTricks

go官方的案例如下，注意一定要先执行rand.Seed(time.Now().UnixNano())

for i := len(a) - 1; i > 0; i-- {
    j := rand.Intn(i + 1)
    a[i], a[j] = a[j], a[i]
}
*/
//返回随机打乱后的切片
func Shuffling(s []interface{}) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回随机打乱后的切片
func ShufflingString(s []string) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingInt(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingInt8(s []int8) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingInt16(s []int16) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingInt32(s []int32) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingInt64(s []int64) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingUint(s []uint) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingUint8(s []uint8) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingUint16(s []uint16) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingUint32(s []uint32) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingUint64(s []uint64) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingUintptr(s []uintptr) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingFloat32(s []float32) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingFloat64(s []float64) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingRune(s []rune) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingByte(s []byte) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingComplex64(s []complex64) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingComplex128(s []complex128) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//返回切片s中是否包含单个元素v
func ShufflingBool(s []bool) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
