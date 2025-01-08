package structure

import (
	"golang.org/x/exp/constraints"
	"time"
)

type NumericalType interface {
	constraints.Integer | constraints.Float
}

// TimeType is a type constraint for time.Time.
type TimeType interface {
	time.Time | *time.Time
}

// Type 约束
type Type interface {
	TimeType | NumericalType | ~string
}

type Iter[T any] []T
