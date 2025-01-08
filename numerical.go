package structure

import (
	"slices"
)

// Sum 求和
func Sum[T NumericalType](nums []T) *T {
	if len(nums) == 0 {
		return nil
	}

	var sum T
	for num := range nums {
		sum += nums[num]
	}

	return &sum
}

// Avg 求平均
func Avg[T NumericalType](nums []T) *T {
	if len(nums) < 1 {
		return nil
	}

	return Address(*Sum(nums) / T(len(nums)))
}

// Max 获取最大值
func Max[T NumericalType](nums []T) *T {
	if len(nums) == 0 {
		return nil
	}

	return Address(slices.Max(nums))
}

// Min 获取最小值
func Min[T NumericalType](nums []T) *T {
	if len(nums) < 1 {
		return nil
	}
	return Address(slices.Min(nums))
}

func Address[T any](data T) *T {
	return &data
}

func Default[T any](data *T) T {
	if data != nil {
		return *data
	}

	var zero T
	return zero
}
