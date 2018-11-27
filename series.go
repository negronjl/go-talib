package talib

import (
	"time"
)

type Series interface {
	Len()	int
	Data(i int)	float64
	VData()	[]float64
}

type TimeSeries interface {
	Len()	int
	Data(time.Time)	float64
	Index(time.Time)	int
}

type TaSeries interface {
	Len() int
	High() []float64
	Open() []float64
	Close() []float64
	Low() []float64
	Volume()	[]float64
}

type SimpleSeries struct {
	Highs   []float64
	Opens   []float64
	Closes  []float64
	Lows    []float64
	Volumes []float64
}

func (s SimpleSeries) Len() int {
	return len(s.Highs)
}

func (s SimpleSeries) High() []float64 {
	return s.Highs
}

func (s SimpleSeries) Open() []float64 {
	return s.Opens
}

func (s SimpleSeries) Close() []float64 {
	return s.Closes
}

func (s SimpleSeries) Low() []float64 {
	return s.Lows
}
