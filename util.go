// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//元素个数超过多少时启用Hash算法，可根据实际情况自行调整该数值
//在切片的去重取交集等算法中，但元素个数较大时，运用hashSet对应的相关算法性能会更优
//注意，启用hashSet后，内存占用将会提高
var HashAlgorithmSwitchNum = 1000

//返回最大值
func maxInt(args ...int) int {
	max := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
	}
	return max
}

//返回最小值
func minInt(args ...int) int {
	min := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return min
}
