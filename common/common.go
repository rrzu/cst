package common

import (
	"fmt"

	"github.com/spf13/cast"
)

// InSlice element 是否存在于切片 slice 里
func InSlice[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

// Round2F 保留两位小数
func Round2F(data float64) float64 {
	return cast.ToFloat64(fmt.Sprintf("%.2f", data))
}

// ToPtr 创建任意类型的指针
func ToPtr[T any](v T) *T {
	return &v
}
