// Copyright 2020 slice Author(https://github.com/yudeguang/slice). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/slice.
package slice

//切片筛选通用函数
/*
//结构体
type Student struct {
	Name   string
	Age    int
	Height int
}

//样本数据
Students := []Student{
	Student{"Danny", 15, 165},
	Student{"Jacky", 16, 180},
	Student{"Alan", 17, 172},
	Student{"Sandy", 18, 168},
}
//循环匹配主函数
func Filter(from []Student, where func(Student) bool) (result []Student) {
	for _, line := range from {
		if where(line) {
			result = append(result, line)
		}
	}
	return result
}
//匹配条件函数
func Where() func(Student) bool {
	return func(s Student) bool {
		return (s.Age > 15) && (s.Height > 170)
	}
}
*/
func Filter(from []interface{}, where func(interface{}) bool) (result []interface{}) {
	for _, line := range from {
		if where(line) {
			result = append(result, line)
		}
	}
	return result
}
