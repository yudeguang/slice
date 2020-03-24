// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotIn(m, n []interface{}) (mNotInN []interface{}) {
	mNotInN = make([]interface{}, 0, len(m))
	for _, v := range m {
		if !Contains(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInString(m, n []string) (mNotInN []string) {
	mNotInN = make([]string, 0, len(m))
	for _, v := range m {
		if !ContainsString(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInInt(m, n []int) (mNotInN []int) {
	mNotInN = make([]int, 0, len(m))
	for _, v := range m {
		if !ContainsInt(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInInt8(m, n []int8) (mNotInN []int8) {
	mNotInN = make([]int8, 0, len(m))
	for _, v := range m {
		if !ContainsInt8(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInInt16(m, n []int16) (mNotInN []int16) {
	mNotInN = make([]int16, 0, len(m))
	for _, v := range m {
		if !ContainsInt16(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInInt32(m, n []int32) (mNotInN []int32) {
	mNotInN = make([]int32, 0, len(m))
	for _, v := range m {
		if !ContainsInt32(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInInt64(m, n []int64) (mNotInN []int64) {
	mNotInN = make([]int64, 0, len(m))
	for _, v := range m {
		if !ContainsInt64(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInUint(m, n []uint) (mNotInN []uint) {
	mNotInN = make([]uint, 0, len(m))
	for _, v := range m {
		if !ContainsUint(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInUint8(m, n []uint8) (mNotInN []uint8) {
	mNotInN = make([]uint8, 0, len(m))
	for _, v := range m {
		if !ContainsUint8(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInUint16(m, n []uint16) (mNotInN []uint16) {
	mNotInN = make([]uint16, 0, len(m))
	for _, v := range m {
		if !ContainsUint16(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInUint32(m, n []uint32) (mNotInN []uint32) {
	mNotInN = make([]uint32, 0, len(m))
	for _, v := range m {
		if !ContainsUint32(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInUint64(m, n []uint64) (mNotInN []uint64) {
	mNotInN = make([]uint64, 0, len(m))
	for _, v := range m {
		if !ContainsUint64(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInUintptr(m, n []uintptr) (mNotInN []uintptr) {
	mNotInN = make([]uintptr, 0, len(m))
	for _, v := range m {
		if !ContainsUintptr(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInFloat32(m, n []float32) (mNotInN []float32) {
	mNotInN = make([]float32, 0, len(m))
	for _, v := range m {
		if !ContainsFloat32(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInFloat64(m, n []float64) (mNotInN []float64) {
	mNotInN = make([]float64, 0, len(m))
	for _, v := range m {
		if !ContainsFloat64(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInRune(m, n []rune) (mNotInN []rune) {
	mNotInN = make([]rune, 0, len(m))
	for _, v := range m {
		if !ContainsRune(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInByte(m, n []byte) (mNotInN []byte) {
	mNotInN = make([]byte, 0, len(m))
	for _, v := range m {
		if !ContainsByte(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInComplex64(m, n []complex64) (mNotInN []complex64) {
	mNotInN = make([]complex64, 0, len(m))
	for _, v := range m {
		if !ContainsComplex64(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInComplex128(m, n []complex128) (mNotInN []complex128) {
	mNotInN = make([]complex128, 0, len(m))
	for _, v := range m {
		if !ContainsComplex128(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}

//返回集合m不在集合n中的所有值，也即m-n，注意在此不对m做去重操作
func NotInBool(m, n []bool) (mNotInN []bool) {
	mNotInN = make([]bool, 0, len(m))
	for _, v := range m {
		if !ContainsBool(n, v) {
			mNotInN = append(mNotInN, v)
		}
	}
	return mNotInN
}
