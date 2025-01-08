package structure

import (
	"fmt"
	"testing"
	"time"
)

type NumericalDate struct {
	Da *time.Time
	Fl float64
	In int
}

func SCNumericalDate(num int) []NumericalDate {
	var res = make([]NumericalDate, num)
	for i := 0; i < num; i++ {
		in := MathRandom[int](1, 10000)
		fl := MathRandom[float64](1, 10000)
		res[i] = NumericalDate{
			Fl: fl,
			In: in,
		}
	}
	return res
}

func TestNumerical(t *testing.T) {
	var list = SCNumericalDate(10)
	slices := StructField[int](list, "In")
	fmt.Println(Sum[int](slices))
	fmt.Println(Avg[int](slices))
	fmt.Println(Max[int](slices))
	fmt.Println(Min[int](slices))

	flt := StructField[float64](list, "Fl")
	fmt.Println(Sum[float64](flt))
	fmt.Println(Avg[float64](flt))
	fmt.Println(Max[float64](flt))
	fmt.Println(Min[float64](flt))

}
