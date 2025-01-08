package structure

import (
	"sort"
	"time"
)

type TimeSort Iter[time.Time]

func (ts TimeSort) Len() int { return len(ts) }
func (ts TimeSort) Less(i, j int) bool {
	return ts[i].Before(ts[j])
}
func (ts TimeSort) Swap(i, j int) { ts[i], ts[j] = ts[j], ts[i] }

func (ts TimeSort) Min() *time.Time {
	if ts == nil || ts.Len() == 0 {
		return nil
	}

	sort.Sort(ts)
	return &ts[0]
}

func (ts TimeSort) Max() *time.Time {
	if ts == nil || ts.Len() == 0 {
		return nil
	}
	sort.Sort(ts)
	return &ts[len(ts)-1]
}
